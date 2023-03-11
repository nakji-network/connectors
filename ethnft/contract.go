package ethnft

import (
	"github.com/nakji-network/connectors/ethnft/smart-contracts/erc1155"
	"github.com/nakji-network/connectors/ethnft/smart-contracts/erc721"

	"github.com/rs/zerolog/log"
)

func GetContracts() []*Contract {
	erc721ABI, err := erc721.IERC721MetaData.GetAbi()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read contract ABI")
	}

	erc1155ABI, err := erc1155.IERC1155MetaData.GetAbi()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read contract ABI")
	}

	return []*Contract{
		{
			ABI:           erc721ABI,
			Name:          "ERC721",
			MessageParser: &erc721.SmartContract{},
		},
		{
			ABI:           erc1155ABI,
			Name:          "ERC1155",
			MessageParser: &erc1155.SmartContract{},
		},
	}
}
