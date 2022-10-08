package near

import (
	"context"

	"github.com/nakji-network/connector"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	ConnectorName string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
	Namespace     string
	WsPort        string
}

type Connector struct {
	*connector.Connector
	*Config
	client *Client
}

func NewConnector(c *connector.Connector, config *Config) *Connector {

	client := NewClient(config)

	return &Connector{
		Connector: c,
		Config:    config,
		client:    client,
	}
}

func (c *Connector) Start() {
	ctx := context.Background()

	c.client.Start(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-c.client.events:
			if msg := c.parse(event); msg != nil {
				c.EventSink <- msg
			}
		}
	}
}

func (c *Connector) parse(event *NearMessage) proto.Message {
	if msg := event.GetBlock(); msg != nil {
		return msg
	} else if msg := event.GetTx(); msg != nil {
		return msg
	}
	return nil
}
