package candymachine

import "google.golang.org/protobuf/proto"

var TopicTypes = map[string]proto.Message{
	"nakji.candymachine.0_0_0.nft_mintnft":                &MintNFT{},
	"nakji.candymachine.0_0_0.nft_updatecandymachine":     &UpdateCandyMachine{},
	"nakji.candymachine.0_0_0.nft_addconfiglines":         &AddConfigLines{},
	"nakji.candymachine.0_0_0.nft_initializecandymachine": &InitializeCandyMachine{},
	"nakji.candymachine.0_0_0.nft_updateauthority":        &UpdateAuthority{},
	"nakji.candymachine.0_0_0.nft_withdrawfunds":          &WithdrawFunds{},
}
