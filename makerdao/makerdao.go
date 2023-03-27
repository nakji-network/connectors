package makerdao

import (
	"context"
	"fmt"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/makerdao/ArbitrumDaiBridge"
	"github.com/nakji-network/connectors/makerdao/ArbitrumGovRelay"
	"github.com/nakji-network/connectors/makerdao/ArbitrumL2Dai"
	"github.com/nakji-network/connectors/makerdao/ArbitrumL2DaiGateway"
	"github.com/nakji-network/connectors/makerdao/BAT"
	"github.com/nakji-network/connectors/makerdao/Bal"
	Comp "github.com/nakji-network/connectors/makerdao/Comp"
	"github.com/nakji-network/connectors/makerdao/CrvV1EthSTEth"
	"github.com/nakji-network/connectors/makerdao/Eth"
	"github.com/nakji-network/connectors/makerdao/GUSD"
	"github.com/nakji-network/connectors/makerdao/GUniV3DaiUSDC2"
	"github.com/nakji-network/connectors/makerdao/IlkRegistry"
	"github.com/nakji-network/connectors/makerdao/JoinFab"
	"github.com/nakji-network/connectors/makerdao/KNC"
	"github.com/nakji-network/connectors/makerdao/LRC"
	"github.com/nakji-network/connectors/makerdao/Link"
	"github.com/nakji-network/connectors/makerdao/MCDCat"
	"github.com/nakji-network/connectors/makerdao/MCDDai"
	"github.com/nakji-network/connectors/makerdao/MCDDog"
	"github.com/nakji-network/connectors/makerdao/MCDESM"
	"github.com/nakji-network/connectors/makerdao/MCDEnd"
	"github.com/nakji-network/connectors/makerdao/MCDFlap"
	"github.com/nakji-network/connectors/makerdao/MCDFlash"
	"github.com/nakji-network/connectors/makerdao/MCDFlop"
	"github.com/nakji-network/connectors/makerdao/MCDGov"
	"github.com/nakji-network/connectors/makerdao/MCDPSMGUSDA"
	"github.com/nakji-network/connectors/makerdao/MCDPSMPaxA"
	"github.com/nakji-network/connectors/makerdao/MCDPSMUSDCA"
	"github.com/nakji-network/connectors/makerdao/MCDVestDai"
	"github.com/nakji-network/connectors/makerdao/MCDVestMkr"
	"github.com/nakji-network/connectors/makerdao/MCDVestMkrTreasury"
	"github.com/nakji-network/connectors/makerdao/MIP21LiquidationOracle"
	"github.com/nakji-network/connectors/makerdao/Mana"
	"github.com/nakji-network/connectors/makerdao/Matic"
	"github.com/nakji-network/connectors/makerdao/OptimismL2Dai"
	"github.com/nakji-network/connectors/makerdao/OptimismL2DaiTokenBridge"
	"github.com/nakji-network/connectors/makerdao/RWA"
	"github.com/nakji-network/connectors/makerdao/RWAConduit"
	"github.com/nakji-network/connectors/makerdao/RWAUrn"
	"github.com/nakji-network/connectors/makerdao/USDT"
	"github.com/nakji-network/connectors/makerdao/Uni"
	"github.com/nakji-network/connectors/makerdao/UniV2"
	"github.com/nakji-network/connectors/makerdao/WBTC"
	"github.com/nakji-network/connectors/makerdao/WSTEth"
	"github.com/nakji-network/connectors/makerdao/YFi"
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
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
	}
}

func (c *Connector) Start() {
	addresses := GetAddresses(ContractAddresses[c.NetworkName])
	c.contracts = GetContracts(ContractAddresses[c.NetworkName])

	ctx := context.Background()

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, addresses, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s connection error", c.NetworkName))
	}

	//	Check whether local file is up-to-date
	// if c.NetworkName == "ethereum" {
	// 	go CheckLatestAddresses(c.ChainClients.Ethereum(ctx), c.ContractsUrl, c.FactoryAddress, ContractAddresses["ethereum"])
	// }

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
				c.EventSink <- msg
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

	time, err := c.sub.GetBlockTime(context.Background(), vLog)
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
	// case "Aave":
	// 	return &Aave.SmartContract{}
	// case "Adai":
	// 	return &Adai.SmartContract{}
	case "ArbitrumDaiBridge":
		return &ArbitrumDaiBridge.SmartContract{}
	// case "ArbitrumEscrow":
	// 	return &ArbitrumEscrow.SmartContract{}
	case "ArbitrumGovRelay":
		return &ArbitrumGovRelay.SmartContract{}
	case "ArbitrumL2Dai":
		return &ArbitrumL2Dai.SmartContract{}
	case "ArbitrumL2DaiGateway":
		return &ArbitrumL2DaiGateway.SmartContract{}
	// case "ArbitrumL2GovRelay":
	// 	return &ArbitrumL2GovRelay.SmartContract{}
	case "Bal":
		return &Bal.SmartContract{}
	case "BAT":
		return &BAT.SmartContract{}
	// case "CdpManager":
	// 	return &CdpManager.SmartContract{}
	// case "CdpRegistry":
	// 	return &CdpRegistry.SmartContract{}
	// case "Changelog":
	// 	return &Changelog.SmartContract{}
	// case "ClipperMom":
	// 	return &ClipperMom.SmartContract{}
	case "Comp":
		return &Comp.SmartContract{}
	case "CrvV1EthSTEth":
		return &CrvV1EthSTEth.SmartContract{}
	// case "DirectMom":
	// 	return &DirectMom.SmartContract{}
	// case "DsrManager":
	// 	return &DsrManager.SmartContract{}
	case "Eth":
		return &Eth.SmartContract{}
	// case "FlipperMom":
	// 	return &FlipperMom.SmartContract{}
	// case "GovGuard":
	// 	return &GovGuard.SmartContract{}
	// case "GUniV3DaiUSDC1":
	// 	return &GUniV3DaiUSDC1.SmartContract{}
	case "GUniV3DaiUSDC2":
		return &GUniV3DaiUSDC2.SmartContract{}
	case "GUSD":
		return &GUSD.SmartContract{}
	case "IlkRegistry":
		return &IlkRegistry.SmartContract{}
	case "JoinFab":
		return &JoinFab.SmartContract{}
	case "KNC":
		return &KNC.SmartContract{}
	// case "LerpFab":
	// 	return &LerpFab.SmartContract{}
	case "Link":
		return &Link.SmartContract{}
	case "LRC":
		return &LRC.SmartContract{}
	case "Mana":
		return &Mana.SmartContract{}
	case "Matic":
		return &Matic.SmartContract{}
	// case "MCDADM":
	// 	return &MCDADM.SmartContract{}
	case "MCDCat":
		return &MCDCat.SmartContract{}
	// case "MCDCropper":
	// 	return &MCDCropper.SmartContract{}
	// case "MCDCropper_IMP":
	// 	return &MCDCropper_IMP.SmartContract{}
	case "MCDDai":
		return &MCDDai.SmartContract{}
	// case "MCDDeploy":
	// 	return &MCDDeploy.SmartContract{}
	case "MCDDog":
		return &MCDDog.SmartContract{}
	case "MCDEnd":
		return &MCDEnd.SmartContract{}
	case "MCDESM":
		return &MCDESM.SmartContract{}
	case "MCDFlap":
		return &MCDFlap.SmartContract{}
	case "MCDFlash":
		return &MCDFlash.SmartContract{}
	case "MCDFlop":
		return &MCDFlop.SmartContract{}
	case "MCDGov":
		return &MCDGov.SmartContract{}
	// case "MCDIamAutoLine":
	// 	return &MCDIamAutoLine.SmartContract{}
	// case "MCDJug":
	// 	return &MCDJug.SmartContract{}
	// case "MCDPause":
	// 	return &MCDPause.SmartContract{}
	// case "MCDPot":
	// 	return &MCDPot.SmartContract{}
	case "MCDPSMGUSDA":
		return &MCDPSMGUSDA.SmartContract{}
	case "MCDPSMPaxA":
		return &MCDPSMPaxA.SmartContract{}
	case "MCDPSMUSDCA":
		return &MCDPSMUSDCA.SmartContract{}
	// case "MCDSpot":
	// 	return &MCDSpot.SmartContract{}
	// case "MCDVat":
	// 	return &MCDVat.SmartContract{}
	case "MCDVestDai":
		return &MCDVestDai.SmartContract{}
	// case "MCDVestDai_LEGACY":
	// 	return &MCDVestDai_LEGACY.SmartContract{}
	case "MCDVestMkr":
		return &MCDVestMkr.SmartContract{}
	case "MCDVestMkrTreasury":
		return &MCDVestMkrTreasury.SmartContract{}
	// case "MCDVow":
	// 	return &MCDVow.SmartContract{}
	case "MIP21LiquidationOracle":
		return &MIP21LiquidationOracle.SmartContract{}
	// case "OptimismDaiBridge":
	// 	return &OptimismDaiBridge.SmartContract{}
	// case "OptimismEscrow":
	// 	return &OptimismEscrow.SmartContract{}
	// case "OptimismGovRelay":
	// 	return &OptimismGovRelay.SmartContract{}
	case "OptimismL2Dai":
		return &OptimismL2Dai.SmartContract{}
	case "OptimismL2DaiTokenBridge":
		return &OptimismL2DaiTokenBridge.SmartContract{}
	// case "OptimismL2GovernanceRelay":
	// 	return &OptimismL2GovernanceRelay.SmartContract{}
	// case "OSMMom":
	// 	return &OSMMom.SmartContract{}
	// case "Pax":
	// 	return &Pax.SmartContract{}
	// case "PaxUSD":
	// 	return &PaxUSD.SmartContract{}
	// case "ProxyDeployer":
	// 	return &ProxyDeployer.SmartContract{}
	// case "ProxyFactory":
	// 	return &ProxyFactory.SmartContract{}
	// case "RenBTC":
	// 	return &RenBTC.SmartContract{}
	case "RWA":
		return &RWA.SmartContract{}
	case "RWAConduit":
		return &RWAConduit.SmartContract{}
	case "RWAUrn":
		return &RWAUrn.SmartContract{}
	// case "STEth":
	// 	return &STEth.SmartContract{}
	// case "TUSD":
	// 	return &TUSD.SmartContract{}
	case "Uni":
		return &Uni.SmartContract{}
	case "UniV2":
		return &UniV2.SmartContract{}
	// case "USDC":
	// 	return &USDC.SmartContract{}
	case "USDT":
		return &USDT.SmartContract{}
	// case "VOTE_DELEGATE_ProxyFactory":
	// 	return &VOTE_DELEGATE_ProxyFactory.SmartContract{}
	// case "VOTE_ProxyFactory":
	// 	return &VOTE_ProxyFactory.SmartContract{}
	case "WBTC":
		return &WBTC.SmartContract{}
	case "WSTEth":
		return &WSTEth.SmartContract{}
	case "YFi":
		return &YFi.SmartContract{}
	case "ZRX":
		return &ZRX.SmartContract{}
	}

	return nil
}
