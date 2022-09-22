package perpetual

import (
	"fmt"
	"strings"

	blepethclient "blep.ai/data/chain/ethereum/ethclient"
	blepcommon "blep.ai/data/common"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L1/PerpRewardVesting"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/Amm"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/AmmReader"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/ChainlinkPriceFeed"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/ClearingHouse"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/InsuranceFund"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/L2PriceFeed"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DbConfig       string
	Env            string
	KafkaUrl       string
	ContractsUrl   string
	FactoryAddress string
	Author         string
	ProtocolName   string
	ConnectorName  string
	NetworkName    string
	Version        string
	RpcUrls        []string
	FromBlock      uint64
	NumBlocks      uint64
}

type Connector struct {
	*Config
	sub         blepethclient.ISubscription
	contracts   map[string]*Contract
	db          *database.Database
	sink        chan *kafkautils.Message
	topicPrefix string
}

func New(config *Config) *Connector {
	return &Connector{
		Config:      config,
		sink:        make(chan *kafkautils.Message),
		topicPrefix: fmt.Sprintf(".fct.%s.%s.%s", config.Author, config.ProtocolName, config.Version),
	}
}

func (c *Connector) Start() {
	c.setup()
	defer c.close()
	c.listen()
}

func (c *Connector) setup() {
	// Load topics
	kafkautils.TopicTypeRegistry.Load(TopicTypes)

	// Connect Database
	db, err := database.New(c.DbConfig)
	if err != nil {
		log.Fatal().Err(err).Str("dsn", c.DbConfig).Msg("Timescaledb connection failed")
	}
	c.db = db

	// Create Kafka sink
	if kp, err := kafkautils.NewProducer(c.KafkaUrl, c.ConnectorName); err == nil {
		kp.EnableTransactions()
		go kp.WriteAndCommitSink(c.sink)
	} else {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	addresses := GetAddresses()
	c.contracts = GetContracts()

	if sub, err := blepethclient.NewSubscription(c.NetworkName, c.RpcUrls, addresses.L2, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Str("network", c.NetworkName).Msg("connection error")
	}

	//	Check whether local file is up-to-date
	if c.NetworkName == "xdai" {
		go CheckLatestAddresses(c.ContractsUrl, c.FactoryAddress, c.sub.ClientPool())
	}
}

func (c *Connector) listen() {
	c.sub.Resubscribe()

	for {
		select {
		case <-c.sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.sub.Err():
			if err != nil {
				log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			}
			c.close()
			return

		//	Listen to event logs
		case vLog := <-c.sub.Logs():
			if msg := c.parse(vLog); msg != nil {
				c.sink <- msg
			}
		}
	}
}

func (c *Connector) close() {
	c.db.Close()
}

func (c *Connector) parse(vLog types.Log) *kafkautils.Message {
	address := vLog.Address.String()
	if c.contracts[address] == nil {
		log.Info().Str("address", address).Msg("Event from unsupported address")
		return nil
	}
	contractAbi := *c.contracts[address].ABI
	contractName := c.contracts[address].Name
	contractType := c.contracts[address].Type

	abiEvent, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Str("contract name", contractName).Err(err).Msg("Failed to get event from ABI")
		return nil
	}

	topic := fmt.Sprintf("%s.%s_%s", c.topicPrefix, strings.ToLower(contractType), strings.ToLower(abiEvent.Name))
	kafkaTopic := kafkautils.MustParseTopic(topic, c.Env)
	kafkaKey := kafkautils.NewKey(c.ConnectorName, vLog.Address.String())

	time, err := c.sub.GetBlockTime(vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := blepcommon.UnixToTimestampPb(int64(time * 1000))

	if smartContract := getContract(contractName); smartContract != nil {
		return &kafkautils.Message{
			Topic:    kafkaTopic,
			Key:      kafkaKey,
			ProtoMsg: smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp),
		}
	}
	return nil
}

func getContract(contractName string) ISmartContract {
	switch contractName {
	case "Amm":
		return &Amm.SmartContract{}
	case "AmmReader":
		return &AmmReader.SmartContract{}
	// case "ChainlinkL1":
	// 	return &ChainlinkL1.SmartContract{}
	case "ChainlinkPriceFeed":
		return &ChainlinkPriceFeed.SmartContract{}
	case "ClearingHouse":
		return &ClearingHouse.SmartContract{}
	// case "ClientBridge":
	// 	return &ClientBridge.SmartContract{}
	// case "FeeRewardPoolL1":
	// 	return &FeeRewardPoolL1.SmartContract{}
	// case "FeeTokenPoolDispatcherL1":
	// 	return &FeeTokenPoolDispatcherL1.SmartContract{}
	case "InsuranceFund":
		return &InsuranceFund.SmartContract{}
	case "L2PriceFeed":
		return &L2PriceFeed.SmartContract{}
	// case "MetaTxGateway":
	// 	return &MetaTxGateway.SmartContract{}
	case "PerpRewardVesting":
		return &PerpRewardVesting.SmartContract{}
		// case "RootBridge":
		// 	return &RootBridge.SmartContract{}
		// case "RootBridgeV2":
		// 	return &RootBridgeV2.SmartContract{}
		// case "StakedPerpToken":
		// 	return &StakedPerpToken.SmartContract{}
	}

	return nil
}
