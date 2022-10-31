package woofi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func GetContracts(addressMap map[string][]string) map[string]*Contract {
	contracts := map[string]*Contract{}
	for contractName, addresses := range addressMap {
		if abiStr, ok := ABIs[contractName]; ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			for _, address := range addresses {
				contracts[address] = &Contract{
					ABI:  &abiObj,
					Name: contractName,
					Type: contractName,
				}
			}
		}
	}

	return contracts
}

func GetAddresses(addressMap map[string][]string) []common.Address {
	addressSlice := make([]common.Address, 0)

	for _, addresses := range addressMap {
		for _, v := range addresses {
			addressSlice = append(addressSlice, common.HexToAddress(v))
		}
	}

	return addressSlice
}
