package solanatoken

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"

	bin "github.com/gagliardetto/binary"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var TopicTypes = map[string]proto.Message{
	"nakji.solanatoken.0_0_0.token_transfer":            &Transfer{},
	"nakji.solanatoken.0_0_0.token_initializemint":      &InitializeMint{},
	"nakji.solanatoken.0_0_0.token_initializeaccount":   &InitializeAccount{},
	"nakji.solanatoken.0_0_0.token_mintto":              &MintTo{},
	"nakji.solanatoken.0_0_0.token_setauthority":        &SetAuthority{},
	"nakji.solanatoken.0_0_0.token_closeaccount":        &CloseAccount{},
	"nakji.solanatoken.0_0_0.token_burn":                &Burn{},
	"nakji.solanatoken.0_0_0.token_approve":             &Approve{},
	"nakji.solanatoken.0_0_0.token_revoke":              &Revoke{},
	"nakji.solanatoken.0_0_0.token_freezeaccount":       &FreezeAccount{},
	"nakji.solanatoken.0_0_0.token_initializemultisig":  &InitializeMultisig{},
	"nakji.solanatoken.0_0_0.token_syncnative":          &SyncNative{},
	"nakji.solanatoken.0_0_0.token_thawaccount":         &ThawAccount{},
	"nakji.solanatoken.0_0_0.token_transferchecked":     &TransferChecked{},
	"nakji.solanatoken.0_0_0.token_approvechecked":      &ApproveChecked{},
	"nakji.solanatoken.0_0_0.token_minttochecked":       &MintToChecked{},
	"nakji.solanatoken.0_0_0.token_burnchecked":         &BurnChecked{},
	"nakji.solanatoken.0_0_0.token_initializeaccount2":  &InitializeAccount2{},
	"nakji.solanatoken.0_0_0.token_initializeaccount3":  &InitializeAccount3{},
	"nakji.solanatoken.0_0_0.token_initializemint2":     &InitializeMint2{},
	"nakji.solanatoken.0_0_0.token_initializemultisig2": &InitializeMultisig2{},
	"nakji.solanatoken.0_0_0.token_nft":                 &NFT{},
	"nakji.solanatoken.0_0_0.token_trade":               &Trade{},
}

func DecodeTransfer(event *token.Transfer, txResult *rpc.GetTransactionResult, tx *solana.Transaction, signature solana.Signature) *Transfer {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	destination := event.Accounts[1].PublicKey.String()
	owner := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	transfer := &Transfer{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Mint:        getMintAddress(txResult, tx, source),
		Source:      source,
		Destination: destination,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(transfer)
	transfer.Id = hashId

	return transfer
}

func DecodeInitializeMint(event *token.InitializeMint, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeMint {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.GetAccounts()
	mint := accounts[0].PublicKey.String()
	rentSysVar := accounts[1].PublicKey.String()

	initializeMint := &InitializeMint{
		Ts:              timestamp,
		Decimals:        uint32(*event.Decimals),
		MintAuthority:   event.MintAuthority.String(),
		FreezeAuthority: toString(event.FreezeAuthority),
		Mint:            mint,
		RentSysVar:      rentSysVar,
		TxSignature:     signature.String(),
	}

	hashId := hashForId(initializeMint)
	initializeMint.Id = hashId

	return initializeMint
}

func DecodeInitializeAccount(event *token.InitializeAccount, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeAccount {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.GetAccounts()
	account := accounts[0].PublicKey.String()
	mint := accounts[1].PublicKey.String()
	owner := accounts[2].PublicKey.String()
	rentSysVar := accounts[3].PublicKey.String()

	initializeAccount := &InitializeAccount{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Account:     account,
		Mint:        mint,
		Owner:       owner,
		RentsysVar:  rentSysVar,
	}

	hashId := hashForId(initializeAccount)
	initializeAccount.Id = hashId

	return initializeAccount
}

func DecodeMintTo(event *token.MintTo, txResult *rpc.GetTransactionResult, signature solana.Signature) *MintTo {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.GetAccounts()
	mint := accounts[0].PublicKey.String()
	destination := accounts[1].PublicKey.String()
	authority := accounts[2].PublicKey.String()
	signers := make([]string, len(event.Signers))
	for i, signer := range event.Signers {
		signers[i] = signer.PublicKey.String()
	}

	mintTo := &MintTo{
		Ts:            timestamp,
		TxSignature:   signature.String(),
		Mint:          mint,
		Amount:        *event.Amount,
		Destination:   destination,
		MintAuthority: authority,
		Signers:       signers,
	}

	hashId := hashForId(mintTo)
	mintTo.Id = hashId

	return mintTo
}

func DecodeSetAuthority(event *token.SetAuthority, txResult *rpc.GetTransactionResult, signature solana.Signature) *SetAuthority {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	authorityType := *event.AuthorityType
	mintOrAccount := event.Accounts[0].PublicKey.String()
	authority := event.Accounts[1].PublicKey.String()

	signers := make([]string, len(event.Accounts[2:]))
	for i, signer := range event.Accounts[2:] {
		signers[i] = signer.PublicKey.String()
	}

	setAuthority := &SetAuthority{
		Ts:            timestamp,
		TxSignature:   signature.String(),
		AuthorityType: authorityTypeToString[bin.BorshEnum(authorityType)],
		Account:       mintOrAccount,
		NewAuthority:  toString(event.NewAuthority),
		Authority:     authority,
		Signers:       signers,
	}

	hashId := hashForId(setAuthority)
	setAuthority.Id = hashId

	return setAuthority
}

func DecodeCloseAccount(event *token.CloseAccount, txResult *rpc.GetTransactionResult, signature solana.Signature) *CloseAccount {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	account := event.Accounts[0].PublicKey.String()
	destination := event.Accounts[1].PublicKey.String()
	owner := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	closeAccount := &CloseAccount{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Account:     account,
		Destination: destination,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(closeAccount)
	closeAccount.Id = hashId

	return closeAccount
}

func DecodeBurn(event *token.Burn, txResult *rpc.GetTransactionResult, signature solana.Signature) *Burn {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	owner := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	burn := &Burn{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Source:      source,
		Mint:        mint,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(burn)
	burn.Id = hashId

	return burn
}

func DecodeApprove(event *token.Approve, txResult *rpc.GetTransactionResult, signature solana.Signature) *Approve {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	delegate := event.Accounts[1].PublicKey.String()
	owner := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	approve := &Approve{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Source:      source,
		Delegate:    delegate,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(approve)
	approve.Id = hashId

	return approve
}

func DecodeRevoke(event *token.Revoke, txResult *rpc.GetTransactionResult, signature solana.Signature) *Revoke {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	owner := event.Accounts[1].PublicKey.String()
	signers := make([]string, len(event.Accounts[2:]))
	for i, signer := range event.Accounts[2:] {
		signers[i] = signer.PublicKey.String()
	}

	revoke := &Revoke{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Source:      source,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(revoke)
	revoke.Id = hashId

	return revoke
}

func DecodeFreezeAccount(event *token.FreezeAccount, txResult *rpc.GetTransactionResult, signature solana.Signature) *FreezeAccount {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	account := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	authority := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	freezeAccount := &FreezeAccount{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Account:     account,
		Mint:        mint,
		Authority:   authority,
		Signers:     signers,
	}

	hashId := hashForId(freezeAccount)
	freezeAccount.Id = hashId

	return freezeAccount
}

func DecodeInitializeMultisig(event *token.InitializeMultisig, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeMultisig {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	account := event.Accounts[0].PublicKey.String()
	rentSysVar := event.Accounts[1].PublicKey.String()
	signers := make([]string, len(event.Accounts[2:]))
	for i, signer := range event.Accounts[2:] {
		signers[i] = signer.PublicKey.String()
	}

	initializeMultisig := &InitializeMultisig{
		Ts:          timestamp,
		TxSignature: signature.String(),
		M:           uint32(*event.M),
		Account:     account,
		RentSysVar:  rentSysVar,
		Signers:     signers,
	}

	hashId := hashForId(initializeMultisig)
	initializeMultisig.Id = hashId

	return initializeMultisig
}

func DecodeSyncNative(event *token.SyncNative, txResult *rpc.GetTransactionResult, signature solana.Signature) *SyncNative {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	tokenAccount := event.AccountMetaSlice.GetAccounts()[0].PublicKey.String()

	syncNative := &SyncNative{
		Ts:           timestamp,
		TxSignature:  signature.String(),
		TokenAccount: tokenAccount,
	}

	hashId := hashForId(syncNative)
	syncNative.Id = hashId

	return syncNative
}

func DecodeThawAccount(event *token.ThawAccount, txResult *rpc.GetTransactionResult, signature solana.Signature) *ThawAccount {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	account := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	authority := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	thawAccount := &ThawAccount{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Account:     account,
		Mint:        mint,
		Authority:   authority,
		Signers:     signers,
	}

	hashId := hashForId(thawAccount)
	thawAccount.Id = hashId

	return thawAccount
}

func DecodeTransferChecked(event *token.TransferChecked, txResult *rpc.GetTransactionResult, signature solana.Signature) *TransferChecked {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	destination := event.Accounts[2].PublicKey.String()
	owner := event.Accounts[3].PublicKey.String()
	signers := make([]string, len(event.Accounts[4:]))
	for i, signer := range event.Accounts[4:] {
		signers[i] = signer.PublicKey.String()
	}

	transferChecked := &TransferChecked{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Decimals:    uint32(*event.Decimals),
		Source:      source,
		Mint:        mint,
		Destination: destination,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(transferChecked)
	transferChecked.Id = hashId

	return transferChecked
}

func DecodeApproveChecked(event *token.ApproveChecked, txResult *rpc.GetTransactionResult, signature solana.Signature) *ApproveChecked {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	delegate := event.Accounts[2].PublicKey.String()
	owner := event.Accounts[3].PublicKey.String()
	signers := make([]string, len(event.Accounts[4:]))
	for i, signer := range event.Accounts[4:] {
		signers[i] = signer.PublicKey.String()
	}

	approveChecked := &ApproveChecked{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Decimals:    uint32(*event.Decimals),
		Source:      source,
		Mint:        mint,
		Delegate:    delegate,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(approveChecked)
	approveChecked.Id = hashId

	return approveChecked
}

func DecodeMintToChecked(event *token.MintToChecked, txResult *rpc.GetTransactionResult, signature solana.Signature) *MintToChecked {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	mint := event.Accounts[0].PublicKey.String()
	destination := event.Accounts[1].PublicKey.String()
	authority := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	mintToChecked := &MintToChecked{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Decimals:    uint32(*event.Decimals),
		Mint:        mint,
		Destination: destination,
		Authority:   authority,
		Signers:     signers,
	}

	hashId := hashForId(mintToChecked)
	mintToChecked.Id = hashId

	return mintToChecked
}

func DecodeBurnChecked(event *token.BurnChecked, txResult *rpc.GetTransactionResult, signature solana.Signature) *BurnChecked {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	source := event.Accounts[0].PublicKey.String()
	mint := event.Accounts[1].PublicKey.String()
	owner := event.Accounts[2].PublicKey.String()
	signers := make([]string, len(event.Accounts[3:]))
	for i, signer := range event.Accounts[3:] {
		signers[i] = signer.PublicKey.String()
	}

	burnChecked := &BurnChecked{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Amount:      *event.Amount,
		Decimals:    uint32(*event.Decimals),
		Source:      source,
		Mint:        mint,
		Owner:       owner,
		Signers:     signers,
	}

	hashId := hashForId(burnChecked)
	burnChecked.Id = hashId

	return burnChecked
}

func DecodeInitializeAccount2(event *token.InitializeAccount2, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeAccount2 {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.AccountMetaSlice.GetAccounts()
	account := accounts[0].PublicKey.String()
	mint := accounts[1].PublicKey.String()
	rentSysVar := accounts[2].PublicKey.String()

	initializeAccount2 := &InitializeAccount2{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Owner:       event.Owner.String(),
		Account:     account,
		Mint:        mint,
		RentSysVar:  rentSysVar,
	}

	hashId := hashForId(initializeAccount2)
	initializeAccount2.Id = hashId

	return initializeAccount2
}

func DecodeInitializeAccount3(event *token.InitializeAccount3, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeAccount3 {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.AccountMetaSlice.GetAccounts()
	account := accounts[0].PublicKey.String()
	mint := accounts[1].PublicKey.String()

	initializeAccount3 := &InitializeAccount3{
		Ts:          timestamp,
		TxSignature: signature.String(),
		Owner:       event.Owner.String(),
		Account:     account,
		Mint:        mint,
	}

	hashId := hashForId(initializeAccount3)
	initializeAccount3.Id = hashId

	return initializeAccount3
}

func DecodeInitializeMultisig2(event *token.InitializeMultisig2, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeMultisig2 {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	account := event.Accounts[0].PublicKey.String()
	signers := make([]string, len(event.Accounts[1:]))
	for i, signer := range event.Accounts[1:] {
		signers[i] = signer.PublicKey.String()
	}

	initializeMultisig2 := &InitializeMultisig2{
		Ts:          timestamp,
		TxSignature: signature.String(),
		M:           uint32(*event.M),
		Account:     account,
		Signers:     signers,
	}

	hashId := hashForId(initializeMultisig2)
	initializeMultisig2.Id = hashId

	return initializeMultisig2
}

func DecodeInitializeMint2(event *token.InitializeMint2, txResult *rpc.GetTransactionResult, signature solana.Signature) *InitializeMint2 {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	accounts := event.AccountMetaSlice.GetAccounts()
	mint := accounts[0].PublicKey.String()

	initializeMint2 := &InitializeMint2{
		Ts:              timestamp,
		TxSignature:     signature.String(),
		Decimals:        uint32(*event.Decimals),
		MintAuthority:   event.MintAuthority.String(),
		FreezeAuthority: toString(event.FreezeAuthority),
		Mint:            mint,
	}

	hashId := hashForId(initializeMint2)
	initializeMint2.Id = hashId

	return initializeMint2
}

func DecodeNFT(mint *token.Mint, meta *token_metadata.Metadata, txResult *rpc.GetTransactionResult) *NFT {
	timestamp := timestamppb.New(txResult.BlockTime.Time())
	var creators []*Creator
	if meta != nil && meta.Data.Creators != nil {
		creators = make([]*Creator, len(*meta.Data.Creators))
		for i, creator := range *meta.Data.Creators {
			creators[i] = &Creator{
				Address:  creator.Address.String(),
				Share:    uint32(creator.Share),
				Verified: creator.Verified,
			}
		}
	}

	var collection string

	if meta.Collection != nil {
		collection = meta.Collection.Key.String()
	}

	nft := &NFT{
		Address:             meta.Mint.String(),
		Ts:                  timestamp,
		MintAuthority:       mint.MintAuthority.String(),
		FreezeAuthority:     toString(mint.FreezeAuthority),
		UpdateAuthority:     meta.UpdateAuthority.String(),
		PrimarySaleHappened: meta.PrimarySaleHappened,
		Royalty:             uint32(meta.Data.SellerFeeBasisPoints),
		Creators:            creators,
		Name:                cleanNulls(meta.Data.Name),
		Symbol:              cleanNulls(meta.Data.Symbol),
		Uri:                 cleanNulls(meta.Data.Uri),
		Collection:          collection,
	}

	return nft
}

func cleanNulls(input string) string {
	result := bytes.Trim([]byte(input), "\x00")
	return string(result)
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

func toString(pubKey *solana.PublicKey) string {
	if pubKey == nil {
		return ""
	}
	return pubKey.String()
}

func getMintAddress(txResult *rpc.GetTransactionResult, tx *solana.Transaction, source string) string {
	var (
		sourceAddressIdx uint16
		mintAddress      string
	)

	for i, publicKey := range tx.Message.AccountKeys {
		if publicKey.String() == source {
			sourceAddressIdx = uint16(i)
			break
		}
	}

	for _, tokenBalance := range txResult.Meta.PostTokenBalances {
		if tokenBalance.AccountIndex == sourceAddressIdx {
			mintAddress = tokenBalance.Mint.String()
			break
		}
	}

	return mintAddress
}
