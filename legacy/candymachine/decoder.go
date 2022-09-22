package candymachine

import (
	"crypto/sha256"
	"encoding/json"
	candymachinev2 "github.com/gagliardetto/metaplex-go/clients/nft-candy-machine-v2"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DecodeMintNFT(impl *candymachinev2.MintNft, txResult *rpc.GetTransactionResult, signature solana.Signature) *MintNFT {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	candyMachineCreator := impl.AccountMetaSlice[1].PublicKey.String()
	payer := impl.AccountMetaSlice[2].PublicKey.String()
	wallet := impl.AccountMetaSlice[3].PublicKey.String()
	metadata := impl.AccountMetaSlice[4].PublicKey.String()
	mint := impl.AccountMetaSlice[5].PublicKey.String()
	mintAuthority := impl.AccountMetaSlice[6].PublicKey.String()
	updateAuthority := impl.AccountMetaSlice[7].PublicKey.String()
	masterEdition := impl.AccountMetaSlice[8].PublicKey.String()
	tokenMetadataProgram := impl.AccountMetaSlice[9].PublicKey.String()
	tokenProgram := impl.AccountMetaSlice[10].PublicKey.String()
	systemProgram := impl.AccountMetaSlice[11].PublicKey.String()
	rent := impl.AccountMetaSlice[12].PublicKey.String()
	clock := impl.AccountMetaSlice[13].PublicKey.String()
	recentBlockHashes := impl.AccountMetaSlice[14].PublicKey.String()
	instructionSysvarAccount := impl.AccountMetaSlice[15].PublicKey.String()

	mintNFT := &MintNFT{
		Ts:                       timestamp,
		TxSignature:              signature.String(),
		CreatorBump:              uint32(*impl.CreatorBump),
		CandyMachine:             candyMachine,
		CandyMachineCreator:      candyMachineCreator,
		Payer:                    payer,
		Wallet:                   wallet,
		Metadata:                 metadata,
		Mint:                     mint,
		MintAuthority:            mintAuthority,
		UpdateAuthority:          updateAuthority,
		MasterEdition:            masterEdition,
		TokenMetadataProgram:     tokenMetadataProgram,
		TokenProgram:             tokenProgram,
		SystemProgram:            systemProgram,
		Rent:                     rent,
		Clock:                    clock,
		RecentBlockhashes:        recentBlockHashes,
		InstructionSysvarAccount: instructionSysvarAccount,
	}

	hashId := hashForId(mintNFT)
	mintNFT.Id = hashId

	return mintNFT
}

func DecodeUpdateCandyMachine(impl *candymachinev2.UpdateCandyMachine, txResult *rpc.GetTransactionResult, signature solana.Signature) *UpdateCandyMachine {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	authority := impl.AccountMetaSlice[1].PublicKey.String()
	wallet := impl.AccountMetaSlice[2].PublicKey.String()

	updateCandyMachine := &UpdateCandyMachine{
		Ts:           timestamp,
		TxSignature:  signature.String(),
		CandyMachine: candyMachine,
		Authority:    authority,
		Wallet:       wallet,
	}

	hashId := hashForId(updateCandyMachine)
	updateCandyMachine.Id = hashId

	return updateCandyMachine
}

func DecodeAddConfigLines(impl *candymachinev2.AddConfigLines, txResult *rpc.GetTransactionResult, signature solana.Signature) *AddConfigLines {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	authority := impl.AccountMetaSlice[1].PublicKey.String()
	index := impl.Index

	var configLines []*ConfigLine
	for _, cl := range *(impl.ConfigLines) {
		configLine := &ConfigLine{Name: cl.Name, Uri: cl.Uri}
		configLines = append(configLines, configLine)
	}

	addConfigLines := &AddConfigLines{
		Ts:           timestamp,
		TxSignature:  signature.String(),
		CandyMachine: candyMachine,
		Authority:    authority,
		Index:        *index,
		ConfigLines:  configLines,
	}

	hashId := hashForId(addConfigLines)
	addConfigLines.Id = hashId

	return addConfigLines
}

func DecodeInitializeCandyMachine(impl *candymachinev2.InitializeCandyMachine, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeCandyMachine {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	wallet := impl.AccountMetaSlice[1].PublicKey.String()
	authority := impl.AccountMetaSlice[2].PublicKey.String()
	payer := impl.AccountMetaSlice[3].PublicKey.String()
	systemProgram := impl.AccountMetaSlice[4].PublicKey.String()
	rent := impl.AccountMetaSlice[5].PublicKey.String()

	cmd := impl.Data

	uuid := ""
	var price uint64 = 0
	symbol := ""
	var sfbp uint16 = 0
	var maxSupply uint64 = 0
	isMutable := false
	ra := false
	var goliveDate int64 = 0
	wlms := ""
	var creators []string
	gk := ""
	es := &EndSettings{}
	hs := &HiddenSettings{}
	var itemsAvailable uint64 = 0

	if cmd != nil {
		uuid = cmd.Uuid
		price = cmd.Price
		symbol = cmd.Symbol
		sfbp = cmd.SellerFeeBasisPoints
		maxSupply = cmd.MaxSupply
		isMutable = cmd.IsMutable
		ra = cmd.RetainAuthority

		if cmd.GoLiveDate != nil {
			goliveDate = *(cmd.GoLiveDate)
		}

		itemsAvailable = cmd.ItemsAvailable

		for _, c := range cmd.Creators {
			creators = append(creators, c.Address.String())
		}

		if cmd.WhitelistMintSettings != nil {
			wlms = cmd.WhitelistMintSettings.Mint.String()
		}

		if cmd.Gatekeeper != nil {
			gk = cmd.Gatekeeper.GatekeeperNetwork.String()
		}

		if cmd.EndSettings != nil {
			es.EndSettingType = uint32(cmd.EndSettings.EndSettingType)
			es.Number = cmd.EndSettings.Number
		}

		if cmd.HiddenSettings != nil {
			hs.Name = cmd.HiddenSettings.Name
			hs.Uri = cmd.HiddenSettings.Uri
		}
	}

	data := &CandyMachineData{
		Uuid:                  uuid,
		Price:                 price,
		Symbol:                symbol,
		SellerFeeBasisPoints:  uint32(sfbp),
		MaxSupply:             maxSupply,
		IsMutable:             isMutable,
		RetainAuthority:       ra,
		GoLiveDate:            uint64(goliveDate),
		EndSettings:           es,
		Creators:              creators,
		HiddenSettings:        hs,
		WhitelistMintSettings: wlms,
		ItemsAvailable:        itemsAvailable,
		GatekeeperConfig:      gk,
	}

	initializeCandyMachine := &InitializeCandyMachine{
		Ts:            timestamp,
		TxSignature:   signature.String(),
		CandyMachine:  candyMachine,
		Authority:     authority,
		Wallet:        wallet,
		Payer:         payer,
		SystemProgram: systemProgram,
		Rent:          rent,
		Data:          data,
	}

	hashId := hashForId(initializeCandyMachine)
	initializeCandyMachine.Id = hashId

	return initializeCandyMachine
}

func DecodeUpdateAuthority(impl *candymachinev2.UpdateAuthority, txResult *rpc.GetTransactionResult, signature solana.Signature) *UpdateAuthority {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	authority := impl.AccountMetaSlice[1].PublicKey.String()
	wallet := impl.AccountMetaSlice[1].PublicKey.String()

	newAuthority := impl.NewAuthority.String()

	updateAuthority := &UpdateAuthority{
		Ts:           timestamp,
		TxSignature:  signature.String(),
		CandyMachine: candyMachine,
		Authority:    authority,
		Wallet:       wallet,
		NewAuthority: newAuthority,
	}

	hashId := hashForId(updateAuthority)
	updateAuthority.Id = hashId

	return updateAuthority
}

func DecodeWithdrawFunds(impl *candymachinev2.WithdrawFunds, txResult *rpc.GetTransactionResult, signature solana.Signature) *WithdrawFunds {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	candyMachine := impl.AccountMetaSlice[0].PublicKey.String()
	authority := impl.AccountMetaSlice[1].PublicKey.String()

	withdrawFunds := &WithdrawFunds{
		Ts:           timestamp,
		TxSignature:  signature.String(),
		CandyMachine: candyMachine,
		Authority:    authority,
	}

	hashId := hashForId(withdrawFunds)
	withdrawFunds.Id = hashId

	return withdrawFunds
}

func hashForId(protoMsg proto.Message) []byte {
	bytes, err := json.Marshal(protoMsg)
	if err != nil {
		log.Fatal().Err(err).Msg("Encoding failed")
	}

	hash := sha256.New()
	hash.Write(bytes)
	h := hash.Sum(nil)

	log.Debug().Hex("Hash", h).Msg("Deterministic hashing for ID")

	return h
}
