package makerdao

import (
	"fmt"
	"strings"

	blepEthclient "blep.ai/data/chain/ethereum/ethclient"
	blepCommon "blep.ai/data/common"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ARBITRUM_DAI_BRIDGE"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ARBITRUM_GOV_RELAY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ARBITRUM_L2_DAI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ARBITRUM_L2_DAI_GATEWAY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ARBITRUM_L2_GOV_RELAY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/BAL"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/BAT"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/COMP"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ETH"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/GUNIV3DAIUSDC2"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/GUSD"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ILK_REGISTRY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/JOIN_FAB"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/KNC"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/LINK"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/LRC"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MANA"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MATIC"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_CAT"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_DAI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_DOG"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_END"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_ESM"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_FLAP"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_FLASH"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_FLOP"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_GOV"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_PSM_GUSD_A"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_PSM_PAX_A"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_PSM_USDC_A"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_VEST_DAI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_VEST_MKR"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MCD_VEST_MKR_TREASURY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/MIP21_LIQUIDATION_ORACLE"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/OPTIMISM_DAI_BRIDGE"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/OPTIMISM_L2_DAI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/OPTIMISM_L2_DAI_TOKEN_BRIDGE"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/OPTIMISM_L2_GOVERNANCE_RELAY"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/RWA"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/RWA_CONDUIT"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/RWA_URN"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/UNI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/UNIV2"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/USDT"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/WBTC"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/WSTETH"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/YFI"
	"blep.ai/data/connectors/source/makerdao/smart-contracts/ZRX"
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
	sub         blepEthclient.ISubscription
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
	kafkautils.TopicTypeRegistry.Load(TopicTypes[c.NetworkName])

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

	addresses := GetAddresses(ContractAddresses[c.NetworkName])
	c.contracts = GetContracts(ContractAddresses[c.NetworkName])

	if sub, err := blepEthclient.NewSubscription(c.NetworkName, c.RpcUrls, addresses, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Str("network", c.NetworkName).Msg("connection error")
	}

	//	Check whether local file is up-to-date
	if c.NetworkName == "ethereum" {
		go CheckLatestAddresses(c.sub.ClientPool(), c.ContractsUrl, c.FactoryAddress, ContractAddresses["ethereum"])
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
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
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

	topic := fmt.Sprintf("%s.%s_%s", c.topicPrefix, strings.ToLower(contractName), strings.ToLower(abiEvent.Name))
	kafkaTopic := kafkautils.MustParseTopic(topic, c.Env)
	kafkaKey := kafkautils.NewKey(c.ConnectorName, vLog.Address.String())

	time, err := c.sub.GetBlockTime(vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := blepCommon.UnixToTimestampPb(int64(time * 1000))

	if smartContract := getContract(contractType); smartContract != nil {
		return &kafkautils.Message{
			Topic:    kafkaTopic,
			Key:      kafkaKey,
			ProtoMsg: smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp),
		}
	}
	return nil
}

func getContract(contractType string) ISmartContract {
	switch contractType {
	// case "AAVE":
	// 	return &AAVE.SmartContract{}
	// case "ADAI":
	// 	return &ADAI.SmartContract{}
	case "ARBITRUM_DAI_BRIDGE":
		return &ARBITRUM_DAI_BRIDGE.SmartContract{}
	// case "ARBITRUM_ESCROW":
	// 	return &ARBITRUM_ESCROW.SmartContract{}
	case "ARBITRUM_GOV_RELAY":
		return &ARBITRUM_GOV_RELAY.SmartContract{}
	case "ARBITRUM_L2_DAI":
		return &ARBITRUM_L2_DAI.SmartContract{}
	case "ARBITRUM_L2_DAI_GATEWAY":
		return &ARBITRUM_L2_DAI_GATEWAY.SmartContract{}
	case "ARBITRUM_L2_GOV_RELAY":
		return &ARBITRUM_L2_GOV_RELAY.SmartContract{}
	case "BAL":
		return &BAL.SmartContract{}
	case "BAT":
		return &BAT.SmartContract{}
	// case "CDP_MANAGER":
	// 	return &CDP_MANAGER.SmartContract{}
	// case "CHANGELOG":
	// 	return &CHANGELOG.SmartContract{}
	// case "CLIPPER_MOM":
	// 	return &CLIPPER_MOM.SmartContract{}
	case "COMP":
		return &COMP.SmartContract{}
	// case "DIRECT_MOM":
	// 	return &DIRECT_MOM.SmartContract{}
	// case "DSR_MANAGER":
	// 	return &DSR_MANAGER.SmartContract{}
	case "ETH":
		return &ETH.SmartContract{}
	// case "FLIPPER_MOM":
	// 	return &FLIPPER_MOM.SmartContract{}
	// case "GOV_GUARD":
	// 	return &GOV_GUARD.SmartContract{}
	// case "GUNIV3DAIUSDC1":
	// 	return &GUNIV3DAIUSDC1.SmartContract{}
	case "GUNIV3DAIUSDC2":
		return &GUNIV3DAIUSDC2.SmartContract{}
	case "GUSD":
		return &GUSD.SmartContract{}
	case "ILK_REGISTRY":
		return &ILK_REGISTRY.SmartContract{}
	case "JOIN_FAB":
		return &JOIN_FAB.SmartContract{}
	case "KNC":
		return &KNC.SmartContract{}
	// case "LERP_FAB":
	// 	return &LERP_FAB.SmartContract{}
	case "LINK":
		return &LINK.SmartContract{}
	case "LRC":
		return &LRC.SmartContract{}
	case "MANA":
		return &MANA.SmartContract{}
	case "MATIC":
		return &MATIC.SmartContract{}
	// case "MCD_ADM":
	// 	return &MCD_ADM.SmartContract{}
	case "MCD_CAT":
		return &MCD_CAT.SmartContract{}
	case "MCD_DAI":
		return &MCD_DAI.SmartContract{}
	// case "MCD_DEPLOY":
	// 	return &MCD_DEPLOY.SmartContract{}
	case "MCD_DOG":
		return &MCD_DOG.SmartContract{}
	case "MCD_END":
		return &MCD_END.SmartContract{}
	case "MCD_ESM":
		return &MCD_ESM.SmartContract{}
	case "MCD_FLAP":
		return &MCD_FLAP.SmartContract{}
	case "MCD_FLASH":
		return &MCD_FLASH.SmartContract{}
	case "MCD_FLOP":
		return &MCD_FLOP.SmartContract{}
	case "MCD_GOV":
		return &MCD_GOV.SmartContract{}
	// case "MCD_IAM_AUTO_LINE":
	// 	return &MCD_IAM_AUTO_LINE.SmartContract{}
	// case "MCD_JUG":
	// 	return &MCD_JUG.SmartContract{}
	// case "MCD_PAUSE":
	// 	return &MCD_PAUSE.SmartContract{}
	// case "MCD_POT":
	// 	return &MCD_POT.SmartContract{}
	case "MCD_PSM_GUSD_A":
		return &MCD_PSM_GUSD_A.SmartContract{}
	case "MCD_PSM_PAX_A":
		return &MCD_PSM_PAX_A.SmartContract{}
	case "MCD_PSM_USDC_A":
		return &MCD_PSM_USDC_A.SmartContract{}
	// case "MCD_SPOT":
	// 	return &MCD_SPOT.SmartContract{}
	// case "MCD_VAT":
	// 	return &MCD_VAT.SmartContract{}
	case "MCD_VEST_DAI":
		return &MCD_VEST_DAI.SmartContract{}
	case "MCD_VEST_MKR":
		return &MCD_VEST_MKR.SmartContract{}
	case "MCD_VEST_MKR_TREASURY":
		return &MCD_VEST_MKR_TREASURY.SmartContract{}
	// case "MCD_VOW":
	// 	return &MCD_VOW.SmartContract{}
	case "MIP21_LIQUIDATION_ORACLE":
		return &MIP21_LIQUIDATION_ORACLE.SmartContract{}
	case "OPTIMISM_DAI_BRIDGE":
		return &OPTIMISM_DAI_BRIDGE.SmartContract{}
	// case "OPTIMISM_ESCROW":
	// 	return &OPTIMISM_ESCROW.SmartContract{}
	// case "OPTIMISM_GOV_RELAY":
	// 	return &OPTIMISM_GOV_RELAY.SmartContract{}
	case "OPTIMISM_L2_DAI":
		return &OPTIMISM_L2_DAI.SmartContract{}
	case "OPTIMISM_L2_DAI_TOKEN_BRIDGE":
		return &OPTIMISM_L2_DAI_TOKEN_BRIDGE.SmartContract{}
	case "OPTIMISM_L2_GOVERNANCE_RELAY":
		return &OPTIMISM_L2_GOVERNANCE_RELAY.SmartContract{}
	// case "OSM_MOM":
	// 	return &OSM_MOM.SmartContract{}
	// case "PAX":
	// 	return &PAX.SmartContract{}
	// case "PAXUSD":
	// 	return &PAXUSD.SmartContract{}
	// case "PROXY_DEPLOYER":
	// 	return &PROXY_DEPLOYER.SmartContract{}
	// case "PROXY_FACTORY":
	// 	return &PROXY_FACTORY.SmartContract{}
	// case "RENBTC":
	// 	return &RENBTC.SmartContract{}
	case "RWA":
		return &RWA.SmartContract{}
	case "RWA_CONDUIT":
		return &RWA_CONDUIT.SmartContract{}
	case "RWA_URN":
		return &RWA_URN.SmartContract{}
	// case "STETH":
	// 	return &STETH.SmartContract{}
	// case "TUSD":
	// 	return &TUSD.SmartContract{}
	case "UNI":
		return &UNI.SmartContract{}
	case "UNIV2":
		return &UNIV2.SmartContract{}
	// case "USDC":
	// 	return &USDC.SmartContract{}
	case "USDT":
		return &USDT.SmartContract{}
	// case "VOTE_DELEGATE_PROXY_FACTORY":
	// 	return &VOTE_DELEGATE_PROXY_FACTORY.SmartContract{}
	// case "VOTE_PROXY_FACTORY":
	// 	return &VOTE_PROXY_FACTORY.SmartContract{}
	case "WBTC":
		return &WBTC.SmartContract{}
	case "WSTETH":
		return &WSTETH.SmartContract{}
	case "YFI":
		return &YFI.SmartContract{}
	case "ZRX":
		return &ZRX.SmartContract{}
	}

	return nil
}
