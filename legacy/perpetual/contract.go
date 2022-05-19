package perpetual

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	blepethclient "blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/InsuranceFund"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage
}

type perpetual struct {
	Layers *layers `json:"layers"`
}

type layers struct {
	Layer1 *layer `json:"layer1"`
	Layer2 *layer `json:"layer2"`
}

type layer struct {
	Contracts map[string]*contract `json:"contracts"`
}

type contract struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Addresses struct {
	L1 []common.Address
	L2 []common.Address
}

type Contract struct {
	ABI  *abi.ABI
	Name string
	Type string
}

func GetContracts() map[string]*Contract {
	obj := loadAddressesFromFile()

	contracts := map[string]*Contract{}
	for k, v := range obj.Layers.Layer1.Contracts {
		abiStr, ok := ABIs[v.Name]
		if ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			contracts[v.Address] = &Contract{
				ABI:  &abiObj,
				Name: k,
				Type: v.Name,
			}
		}
	}

	for k, v := range obj.Layers.Layer2.Contracts {
		abiStr, ok := ABIs[v.Name]
		if ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			contracts[v.Address] = &Contract{
				ABI:  &abiObj,
				Name: k,
				Type: v.Name,
			}
		}
	}

	return contracts
}

func GetAddresses() *Addresses {
	obj := loadAddressesFromFile()

	addresses := new(Addresses)
	addressesL1 := make([]common.Address, len(obj.Layers.Layer1.Contracts))
	i := 0
	for _, v := range obj.Layers.Layer1.Contracts {
		addressesL1[i] = common.HexToAddress(v.Address)
		i++
	}
	addresses.L1 = addressesL1

	addressesL2 := make([]common.Address, len(obj.Layers.Layer2.Contracts))
	j := 0
	for _, v := range obj.Layers.Layer2.Contracts {
		addressesL2[j] = common.HexToAddress(v.Address)
		j++
	}
	addresses.L2 = addressesL2

	return addresses
}

func CheckLatestAddresses(url string, factoryAddress string, clientPool *blepethclient.ClientPool) {
	local := loadAddressesFromFile()
	dynamicAddresses, err := getDynamicAdresses(factoryAddress, clientPool)
	if err != nil {
		return
	}
	web, err := loadAddressesFromWeb(url)
	if err != nil {
		return
	}

	l1 := web.Layers.Layer1.Contracts
	for k, v := range l1 {
		localValue := local.Layers.Layer1.Contracts[k]
		if localValue == nil || localValue.Address != v.Address {
			log.Warn().Str("contract", k).Str("name", v.Name).Str("address", v.Address).Msg("New contract found on perpetual web")
		}
	}

	l2 := web.Layers.Layer2.Contracts
	for k, v := range l2 {
		localValue := local.Layers.Layer2.Contracts[k]
		if localValue == nil || localValue.Address != v.Address {
			log.Warn().Str("contract", k).Str("name", v.Name).Str("address", v.Address).Msg("New contract found on perpetual web")
		}
	}

	for _, v := range local.Layers.Layer2.Contracts {
		delete(dynamicAddresses, v.Address)
	}

	if len(dynamicAddresses) > 0 {
		for k, _ := range dynamicAddresses {
			log.Warn().Str("address", k).Msg("New contract address found at perpetual factory contract")
		}
	}
}

func CheckAndAddNewAddress(subscription *blepethclient.Subscription, contracts map[string]*Contract, protoMsg *kafkautils.Message) {
	w := strings.Split(protoMsg.Topic.String(), ".")
	if w[5] == "insurancefund" && w[6] == "tokenadded" {
		address := strings.Split(protoMsg.Key.String(), ".")[1]

		for a := range contracts {
			if a == address {
				log.Info().Str("address", address).Msg("Address already exists")
				return
			}
		}

		for _, c := range contracts {
			if c.Type == "Amm" {
				contracts[address] = &Contract{
					ABI:  c.ABI,
					Type: "Amm",
				}
				break
			}
		}

		query := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(address)},
		}

		sub, err := subscription.ClientPool().SubscribeFilterLogs(context.Background(), query, subscription.Logs())
		if err != nil {
			log.Warn().Err(err).Str("address", address).Msg("Failed to subscribe to new address")
		}

		if err = <-sub.Err(); err != nil {
			sub.Unsubscribe()
		}
	}
}

func getDynamicAdresses(factoryAddress string, clientPool *blepethclient.ClientPool) (map[string]struct{}, error) {
	address := common.HexToAddress(factoryAddress)
	client, err := clientPool.RandClient(false)
	if err != nil {
		log.Warn().Err(err).Msg("failed to find available client")
		return nil, err
	}

	caller, err := InsuranceFund.NewInsuranceFundCaller(address, client)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to create caller")
		return nil, err
	}

	addresses, err := caller.GetAllAmms(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get AMM addresses")
		return nil, err
	}

	addressMap := make(map[string]struct{}, len(addresses))
	for _, item := range addresses {
		addressMap[item.String()] = struct{}{}
	}
	return addressMap, nil
}

func loadAddressesFromFile() *perpetual {
	obj := new(perpetual)
	err := json.Unmarshal([]byte(ContractAddresses), &obj)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal file")
	}

	return obj
}

func loadAddressesFromWeb(url string) (*perpetual, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to connect to contracts URL-> " + url)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to read response body")
		return nil, err
	}

	obj := new(perpetual)
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to unmarshal file")
		return nil, err
	}

	return obj, nil
}
