package starknet

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage
}

type Contract struct {
	ABI  *abi.ABI
	Name string
}

func GetContracts(addresses map[string]string) map[string]*Contract {
	contracts := map[string]*Contract{}
	for k, v := range addresses {
		abiStr, ok := ABIs[k]
		if ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			contracts[v] = &Contract{
				ABI:  &abiObj,
				Name: k,
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

func CheckLatestAddresses(url string, addresses map[string]string) {
	web := loadAddressesFromWeb(url)

	for k, v := range web {
		adr, ok := addresses[k]
		if !ok || adr != v {
			log.Warn().Str("contract", k).Str("address", v).Msg("New contract found on starknet web")
		}
	}
}

func loadAddressesFromWeb(url string) map[string]string {
	res, err := http.Get(url)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to connect to contracts URL-> " + url)
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
