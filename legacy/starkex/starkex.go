package starkex

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type StarkexConnector struct {
	KP               kafkautils.ProducerInterface
	DB               *database.Database
	Client           *ethclient.Client
	PositionsTopic   kafkautils.Topic
	Web3HttpProvider string
	Interval         int
	VaultCache       map[string]bool
	InitialBlock     uint64
}

func NewConnector(
	kp kafkautils.ProducerInterface,
	db *database.Database,
	client *ethclient.Client,
	positionsTopic kafkautils.Topic,
	web3httpprovider string,
	interval int,
	initialBlock uint64,
) *StarkexConnector {
	vaultCache := make(map[string]bool)

	return &StarkexConnector{
		KP:               kp,
		DB:               db,
		Client:           client,
		PositionsTopic:   positionsTopic,
		Web3HttpProvider: web3httpprovider,
		Interval:         interval,
		VaultCache:       vaultCache,
		InitialBlock:     initialBlock,
	}
}

func (c *StarkexConnector) Start(ctx context.Context, conf *viper.Viper) {
	workingDir := conf.GetString("starkex.workingdir")

	if workingDir == "" {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal().Msg("no caller information")
		}

		workingDir = filepath.Dir(filename)
	}

	err := os.Chdir(workingDir)
	if err != nil {
		log.Fatal().Str("directory", workingDir).Msg("Could not change directory")
	}

	err = c.KP.EnableTransactions()
	if err != nil {
		log.Fatal().Err(err).Msg("Transaction was not enabled")
	}

	for {
		c.runLoop(ctx, workingDir, conf)
	}
}

const (
	Namespace = "ethereum"
	App       = "dydx"
)

func (c *StarkexConnector) runLoop(ctx context.Context, workingDir string, conf *viper.Viper) {
	fromBlock, toBlock := c.waitForMoreBlocks(ctx)
	c.queryPerpPositions(ctx, conf, workingDir, fromBlock, toBlock)
	outputPath := conf.GetString("starkex.outputPath")
	c.writeToKafka(filepath.Join(workingDir, outputPath))
	c.DB.UpdatelastBlock(Namespace, App, toBlock)
}

func (c *StarkexConnector) waitForMoreBlocks(ctx context.Context) (fromBlock, toBlock uint64) {
	for {
		lastBlock, err := c.Client.BlockNumber(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get last block")
		}

		defaultBlock := big.NewInt(int64(c.InitialBlock))
		fromBlock := c.DB.GetLastBlockQueried(Namespace, App, defaultBlock)
		toBlock := fromBlock.Uint64() + 1000
		if toBlock > lastBlock {
			toBlock = lastBlock
		}

		if toBlock-fromBlock.Uint64() < 100 {
			log.Info().Uint64("blocksForQuerying", toBlock-fromBlock.Uint64()).Msg("wait 1hr for more blocks")
			time.Sleep(time.Duration(1 * time.Hour))
		} else {
			return fromBlock.Uint64(), toBlock
		}
	}
}

func (c *StarkexConnector) queryPerpPositions(ctx context.Context, conf *viper.Viper, workingDir string, fromBlock, toBlock uint64) {
	outputPath := conf.GetString("starkex.outputPath")
	outputFile := filepath.Join(workingDir, outputPath)

	pythonCmd := append(
		conf.GetStringSlice("starkex.pythonCmd"),
		fmt.Sprintf("--from_block=%d", fromBlock),
		fmt.Sprintf("--to_block=%d", toBlock),
		fmt.Sprintf("--web3_http_provider=%s", c.Web3HttpProvider),
		fmt.Sprintf("--output_file=%s", outputFile),
	)

	cmd := exec.Command(pythonCmd[0], pythonCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal().Str("cmd", cmd.String()).Err(err).Msg("failed running python script for perpetual onchain data")
	}
}

type Vault struct {
	Key         string           `json:"stark_key"`
	Assets      map[string]Asset `json:"assets"`
	BlockNumber uint64           `json:"block_num"`
	Ts          uint64           `json:"timestamp"`
}

type Asset struct {
	Type                 string `json:"asset_type"`
	Amount               string `json:"amount"`
	AdditionalAttributes Attr   `json:"additional_attributes"`
}

type Attr struct {
	CachedFundingIndex int64 `json:"cached_funding_index"`
}

func (c *StarkexConnector) writeToKafka(outputPath string) {
	log.Info().Msg("reading output.json")
	content, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read file")
	}

	var vaults []Vault
	json.Unmarshal(content, &vaults)

	for _, vault := range vaults {
		c.insertVault(vault)
		c.publishPosition(vault)
	}

	if len(vaults) > 0 {
		c.DB.RefreshLatestPositions()
	}

	if err := ioutil.WriteFile(outputPath, []byte("[]"), 0644); err != nil {
		log.Fatal().Err(err).Msg("failed to clear output.json")
	}
}

func (c *StarkexConnector) publishPosition(vault Vault) {
	for assetKey, asset := range vault.Assets {
		var position Position
		position.UnmarshalPosition(vault, assetKey, asset)
		transactionKey := []byte("starkex." + vault.Key)
		err := c.KP.WriteAndCommit(c.PositionsTopic, transactionKey, &position)
		if err != nil {
			log.Error().Err(err).Msg("Starkex: write Position to Kafka")
		}
	}
}

func (c *StarkexConnector) insertVault(vault Vault) {
	if _, ok := c.VaultCache[vault.Key]; ok {
		log.Info().Str("starkKey", vault.Key).Msg("Vault already inserted and cached")
		return
	}

	sql, args, err :=
		sq.Insert(`"nakji.starkex.0_0_0.starkex_vault"`).
			Suffix("ON CONFLICT DO NOTHING").
			Columns(`"starkKey"`).
			PlaceholderFormat(sq.Dollar).
			Values(vault.Key).
			ToSql()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to prepare sql insert for vault")
	}

	_, err = c.DB.Exec(context.Background(), sql, args...)

	if err != nil {
		log.Fatal().Err(err).Str("sql", sql).Interface("vals", args).Msg("Insert Vault error")
	}

	c.VaultCache[vault.Key] = true
	log.Info().Str("starkKey", vault.Key).Msg("Vault inserted")
}
