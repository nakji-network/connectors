package mai

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

func GetContracts(addresses map[string]string) map[string]*Contract {
	contracts := map[string]*Contract{}
	for k, v := range addresses {
		if abiStr, ok := ABIs[k]; ok {
			abiObj, err := abi.JSON(strings.NewReader(abiStr))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read contract ABI")
			}

			contracts[v] = &Contract{
				ABI:           &abiObj,
				Name:          k,
				MessageParser: Parsers[k],
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
