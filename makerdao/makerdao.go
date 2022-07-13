package makerdao

import (
	"context"
	"fmt"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/makerdao/ARBITRUM_DAI_BRIDGE"
	"github.com/nakji-network/connectors/makerdao/ARBITRUM_GOV_RELAY"
	"github.com/nakji-network/connectors/makerdao/ARBITRUM_L2_DAI"
	"github.com/nakji-network/connectors/makerdao/ARBITRUM_L2_DAI_GATEWAY"
	"github.com/nakji-network/connectors/makerdao/ARBITRUM_L2_GOV_RELAY"
	"github.com/nakji-network/connectors/makerdao/BAL"
	"github.com/nakji-network/connectors/makerdao/BAT"
	"github.com/nakji-network/connectors/makerdao/COMP"
	"github.com/nakji-network/connectors/makerdao/CRVV1ETHSTETH"
	"github.com/nakji-network/connectors/makerdao/ETH"
	"github.com/nakji-network/connectors/makerdao/GUNIV3DAIUSDC2"
	"github.com/nakji-network/connectors/makerdao/GUSD"
	"github.com/nakji-network/connectors/makerdao/ILK_REGISTRY"
	"github.com/nakji-network/connectors/makerdao/JOIN_FAB"
	"github.com/nakji-network/connectors/makerdao/KNC"
	"github.com/nakji-network/connectors/makerdao/LINK"
	"github.com/nakji-network/connectors/makerdao/LRC"
	"github.com/nakji-network/connectors/makerdao/MANA"
	"github.com/nakji-network/connectors/makerdao/MATIC"
	"github.com/nakji-network/connectors/makerdao/MCD_CAT"
	"github.com/nakji-network/connectors/makerdao/MCD_DAI"
	"github.com/nakji-network/connectors/makerdao/MCD_DOG"
	"github.com/nakji-network/connectors/makerdao/MCD_END"
	"github.com/nakji-network/connectors/makerdao/MCD_ESM"
	"github.com/nakji-network/connectors/makerdao/MCD_FLAP"
	"github.com/nakji-network/connectors/makerdao/MCD_FLASH"
	"github.com/nakji-network/connectors/makerdao/MCD_FLOP"
	"github.com/nakji-network/connectors/makerdao/MCD_GOV"
	"github.com/nakji-network/connectors/makerdao/MCD_PSM_GUSD_A"
	"github.com/nakji-network/connectors/makerdao/MCD_PSM_PAX_A"
	"github.com/nakji-network/connectors/makerdao/MCD_PSM_USDC_A"
	"github.com/nakji-network/connectors/makerdao/MCD_VEST_DAI"
	"github.com/nakji-network/connectors/makerdao/MCD_VEST_MKR"
	"github.com/nakji-network/connectors/makerdao/MCD_VEST_MKR_TREASURY"
	"github.com/nakji-network/connectors/makerdao/MIP21_LIQUIDATION_ORACLE"
	"github.com/nakji-network/connectors/makerdao/OPTIMISM_DAI_BRIDGE"
	"github.com/nakji-network/connectors/makerdao/OPTIMISM_L2_DAI"
	"github.com/nakji-network/connectors/makerdao/OPTIMISM_L2_DAI_TOKEN_BRIDGE"
	"github.com/nakji-network/connectors/makerdao/OPTIMISM_L2_GOVERNANCE_RELAY"
	"github.com/nakji-network/connectors/makerdao/RWA"
	"github.com/nakji-network/connectors/makerdao/RWA_CONDUIT"
	"github.com/nakji-network/connectors/makerdao/RWA_URN"
	"github.com/nakji-network/connectors/makerdao/UNI"
	"github.com/nakji-network/connectors/makerdao/UNIV2"
	"github.com/nakji-network/connectors/makerdao/USDT"
	"github.com/nakji-network/connectors/makerdao/WBTC"
	"github.com/nakji-network/connectors/makerdao/WSTETH"
	"github.com/nakji-network/connectors/makerdao/YFI"
	"github.com/nakji-network/connectors/makerdao/ZRX"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	ContractsUrl   string
	FactoryAddress string
	ConnectorName  string
	NetworkName    string
	FromBlock      uint64
	NumBlocks      uint64
}

type Connector struct {
	*connector.Connector
	*Config
	sub       connector.ISubscription
	contracts map[string]*Contract
	sink      chan protoreflect.ProtoMessage
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
		sink:      make(chan protoreflect.ProtoMessage),
	}
}

func (c *Connector) Start() {
	c.setup()
	go c.InitProduceChannel(c.sink)
	c.listen()
}

func (c *Connector) setup() {
	addresses := GetAddresses(ContractAddresses[c.NetworkName])
	c.contracts = GetContracts(ContractAddresses[c.NetworkName])

	ctx := context.Background()

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, addresses, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s connection error", c.NetworkName))
	}

	//	Check whether local file is up-to-date
	if c.NetworkName == "ethereum" {
		go CheckLatestAddresses(c.ChainClients.Ethereum(ctx), c.ContractsUrl, c.FactoryAddress, ContractAddresses["ethereum"])
	}
}

func (c *Connector) listen() {
	c.sub.Subscribe()

	for {
		select {
		case <-c.sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.sub.Logs():
			if msg := c.parse(vLog); msg != nil {
				c.sink <- msg
			}
		}
	}
}

func (c *Connector) parse(vLog types.Log) protoreflect.ProtoMessage {
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

	time, err := c.sub.GetBlockTime(vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := common.UnixToTimestampPb(int64(time * 1000))

	if smartContract := getContract(contractType); smartContract != nil {
		return smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp)
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
	// case "CDP_REGISTRY":
	// 	return &CDP_REGISTRY.SmartContract{}
	// case "CHANGELOG":
	// 	return &CHANGELOG.SmartContract{}
	// case "CLIPPER_MOM":
	// 	return &CLIPPER_MOM.SmartContract{}
	case "COMP":
		return &COMP.SmartContract{}
	case "CRVV1ETHSTETH":
		return &CRVV1ETHSTETH.SmartContract{}
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
	// case "MCD_CROPPER":
	// 	return &MCD_CROPPER.SmartContract{}
	// case "MCD_CROPPER_IMP":
	// 	return &MCD_CROPPER_IMP.SmartContract{}
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
	// case "MCD_VEST_DAI_LEGACY":
	// 	return &MCD_VEST_DAI_LEGACY.SmartContract{}
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
