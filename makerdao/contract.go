package makerdao

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/nakji-network/connectors/makerdao/ILK_REGISTRY"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage
}

type Contract struct {
	ABI  *abi.ABI
	Name string
	Type string
}

func GetContracts(addresses map[string]string) map[string]*Contract {
	contracts := map[string]*Contract{}
	for k, v := range addresses {
		var t string
		if strings.HasPrefix(k, "PIP_") {
			t = "PIP"
		} else if strings.HasPrefix(k, "MCD_JOIN_") {
			t = "MCD_JOIN"
		} else if strings.HasPrefix(k, "MCD_CLIP_CALC") {
			t = "MCD_CLIP_CALC"
		} else if strings.HasPrefix(k, "MCD_CLIP") {
			t = "MCD_CLIP"
		} else {
			t = k
		}

		//	Ignore the events above for now (e.g. PIP, MCD_JOIN)
		if t != k {
			continue
		}

		if strings.HasPrefix(k, "RWA") {
			if strings.HasSuffix(k, "URN") {
				t = "RWA_URN"
			} else if strings.HasSuffix(k, "CONDUIT") {
				t = "RWA_CONDUIT"
			} else {
				t = "RWA"
			}
		} else if strings.HasPrefix(k, "UNIV2") {
			t = "UNIV2"
		}

		if abiStr, ok := ABIs[t]; ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			contracts[v] = &Contract{
				ABI:  &abiObj,
				Name: t,
				Type: t,
			}
		}
	}

	return contracts
}

func GetAddresses(addresses map[string]string) []common.Address {
	addressSlice := make([]common.Address, len(addresses))
	i := 0
	for _, v := range addresses {
		addressSlice[i] = common.HexToAddress(v)
		i++
	}

	return addressSlice
}

func CheckLatestAddresses(client *ethclient.Client, contractsUrl string, factoryAddress string, addresses map[string]string) {
	dynamicAddresses := getDynamicAdresses(client, factoryAddress)
	web := loadAddressesFromWeb(contractsUrl)

	for k, v := range web {
		adr, ok := addresses[k]
		if !ok || adr != v {
			log.Warn().Str("contract", k).Str("address", v).Msg("New contract found on makerdao web")
		}
	}

	for _, v := range addresses {
		delete(dynamicAddresses, v)
	}

	if len(dynamicAddresses) > 0 {
		for k, v := range dynamicAddresses {
			log.Warn().Str("name", string(v[:])).Str("address", k).Msg("New contract address found at makerdao factory contract")
		}
	}
}

//	This method retrieves vault names for all collateral types (i.e. tokens) from the factory contract.
//	Then it gets different type of contracts for each collateral
func getDynamicAdresses(client *ethclient.Client, factoryAddress string) map[string]string {
	caller, err := ILK_REGISTRY.NewILKREGISTRYCaller(common.HexToAddress(factoryAddress), client)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create caller")
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	names, err := caller.List(opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get dynamic addresses")
	}

	addressMap := make(map[string]string)
	for _, name := range names {
		n := string(name[:])

		//	Collateral tokens
		addressGem := getAddressFromChain(opts, name, caller.Gem)
		addressMap[addressGem] = n

		//	This contract provides price info of the collateral token
		addressPip := getAddressFromChain(opts, name, caller.Pip)
		addressMap[addressPip] = n

		//	Liquidation of collateral
		addressClip := getAddressFromChain(opts, name, caller.Xlip)
		addressMap[addressClip] = n
	}

	return addressMap
}

func getAddressFromChain(opts *bind.CallOpts, name [32]byte, f func(*bind.CallOpts, [32]byte) (common.Address, error)) string {
	address, err := f(opts, name)
	if err != nil {
		log.Warn().Err(err).Str("name", string(name[:])).Msg("Failed to get dynamic address")
		return ""
	}
	return address.String()
}

func loadAddressesFromWeb(contractsUrl string) map[string]string {
	res, err := http.Get(contractsUrl)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to connect to contracts URL-> " + contractsUrl)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to read response body")
	}

	var obj map[string]string
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal file")
	}

	return obj
}
