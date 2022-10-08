package near

import (
	"context"
	"embed"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

//go:embed target/release/nakji_near_client
var f embed.FS

type Client struct {
	events chan *NearMessage
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{
		events: make(chan *NearMessage),
		config: config,
	}
}

func (c *Client) Start(ctx context.Context) {
	c.connect(ctx)
}

func (c *Client) connect(ctx context.Context) {
	bin, err := f.ReadFile("target/release/nakji_near_client")

	if err != nil {
		log.Fatal().Err(err).Msg("failed read embedded Nakji Near Client binary")
	}

	err = os.WriteFile("nakji_near_client", bin, 0755)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to write Nakji Near Client binary to local fs")
	}

	cmd := []string{
		"./nakji_near_client",
	}

	cmd = append(cmd, c.config.WsPort)
	if c.config.FromBlock > 0 {
		cmd = append(cmd, fmt.Sprintf("from-block %d", c.config.FromBlock))
	} else {
		cmd = append(cmd, "from-latest")
	}
	execCmd := exec.CommandContext(ctx, cmd[0], cmd[1:]...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	go func() {
		if err := execCmd.Run(); err != nil {
			log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error running NEAR Rust binary: ", err))
		}
	}()
	time.Sleep(5 * time.Second)

	go c.wsStream(ctx, c.config.WsPort)
}

func (c *Client) wsStream(ctx context.Context, port string) {
	host := "localhost:"
	host = fmt.Sprintf("%s%s", host, port)
	u := url.URL{Scheme: "ws", Host: host}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error connecting to websocket", err))
	}
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("connector shutdown")
			return
		default:
			_, wsMessage, err := conn.ReadMessage()
			if err != nil {
				log.Error().Err(err).Msg(fmt.Sprintf("%s Error reading message from websocket: ", err))
			}

			msg := &NearMessage{}

			if err := proto.Unmarshal(wsMessage, msg); err != nil {
				log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error unmarshalling message: ", err))
			}

			c.events <- msg
		}
	}
}
