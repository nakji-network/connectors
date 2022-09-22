package compound

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	. "blep.ai/data/connectors/source/compound/ctoken"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var topics map[string]kafkautils.Topic

var ts uint64 = 1625646821

const (
	mintDataRaw            = "0x000000000000000000000000f55d7a2f553be0beaedce903103a2a13e9b5508c000000000000000000000000000000000000000000009c3f097d548f03700000000000000000000000000000000000000000000000000000000cb92dff5b5f80"
	redeemDataRaw          = "0x000000000000000000000000d25f28bd5feffe05b832cef7d32f27c7e2225e1100000000000000000000000000000000000000000000000649111a7bd08ca1910000000000000000000000000000000000000000000000000000008307646d85"
	borrowDataRaw          = "0x000000000000000000000000f55d7a2f553be0beaedce903103a2a13e9b5508c000000000000000000000000000000000000000000009c3f097d548f03700000000000000000000000000000000000000000000000009c3f097d548f037000000000000000000000000000000000000000000000001115a0f9094ed17aa1205b"
	repayborrowDataRaw     = "0x0000000000000000000000004e390a00a8392137e6a330f900a33e943c3e5c0a0000000000000000000000004e390a00a8392137e6a330f900a33e943c3e5c0a0000000000000000000000000000000000000000000000000ef1b2dc6377d2cb000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010599784d29e541d88bf81"
	liquidateborrowDataRaw = ""

	mintTopic0            = "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"
	redeemTopic0          = "0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929"
	borrowTopic0          = "0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80"
	repayborrowTopic0     = "0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1"
	liquidateborrowTopic0 = "0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52"
)

func setup() {
	kafkautils.TopicTypeRegistry.Load(TopicTypes)
	topics = map[string]kafkautils.Topic{
		"mint":            kafkautils.MustParseTopic(".fct.nakji.compound.0_0_0.compound_mint", "test"),
		"redeem":          kafkautils.MustParseTopic(".fct.nakji.compound.0_0_0.compound_redeem", "test"),
		"borrow":          kafkautils.MustParseTopic(".fct.nakji.compound.0_0_0.compound_borrow", "test"),
		"repayborrow":     kafkautils.MustParseTopic(".fct.nakji.compound.0_0_0.compound_repayborrow", "test"),
		"liquidateborrow": kafkautils.MustParseTopic(".fct.nakji.compound.0_0_0.compound_liquidateborrow", "test"),
	}
}

func TestWriteEventToChan(t *testing.T) {
	setup()

	contractAbi1, _ := abi.JSON(strings.NewReader(CompoundABI))
	contractAbi2, _ := abi.JSON(strings.NewReader(CompoundABI))
	contractAbis := []abi.ABI{contractAbi1, contractAbi2}

	out := make(chan *kafkautils.Message, 10)

	c := &Connector{Topics: topics}

	mintTopic := buildTopics(mintTopic0)
	mintData := buildData(mintDataRaw)

	redeemTopic := buildTopics(redeemTopic0)
	redeemData := buildData(redeemDataRaw)

	borrowTopic := buildTopics(borrowTopic0)
	borrowData := buildData(borrowDataRaw)

	repayborrowTopic := buildTopics(repayborrowTopic0)
	repayborrowData := buildData(repayborrowDataRaw)

	type args struct {
		evLog ethtypes.Log
	}
	tests := []struct {
		name string
		args args
	}{
		{"nakji.compound.0_0_0.compound_mint",
			args{ethtypes.Log{Topics: mintTopic, Data: mintData}},
		},
		{"nakji.compound.0_0_0.compound_redeem",
			args{ethtypes.Log{Topics: redeemTopic, Data: redeemData}},
		},
		{"nakji.compound.0_0_0.compound_borrow",
			args{ethtypes.Log{Topics: borrowTopic, Data: borrowData}},
		},
		{"nakji.compound.0_0_0.compound_repayborrow",
			args{ethtypes.Log{Topics: repayborrowTopic, Data: repayborrowData}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.WriteEventToChan(contractAbis, tt.args.evLog, ts, out)
			msg := <-out
			var res bool
			switch tt.name {
			case "nakji.compound.0_0_0.compound_mint":
				res = checkMint(msg)
			case "nakji.compound.0_0_0.compound_redeem":
				res = checkRedeem(msg)
			case "nakji.compound.0_0_0.compound_borrow":
				res = checkBorrow(msg)
			case "nakji.compound.0_0_0.compound_repayborrow":
				res = checkRepayBorrow(msg)
			}

			if !res {
				t.Errorf("WriteEventToChan() failed for event %v", tt.name)
			}
			if msg.Topic.Env != "test" {
				t.Errorf("Topic.Env wrong for event %v", tt.name)
			}
			if msg.Topic.Schema() != tt.name {
				t.Errorf("Topic.Schema wrong for event %v", tt.name)
			}
		})
	}
}

func checkMint(message *kafkautils.Message) bool {
	mint := reflect.ValueOf(message.ProtoMsg).Interface().(*Mint)
	minter := []byte{245, 93, 122, 47, 85, 59, 224, 190, 174, 220, 233, 3, 16, 58, 42, 19, 233, 181, 80, 140}
	amount := []byte{156, 63, 9, 125, 84, 143, 3, 112, 0, 0}
	tokens := []byte{12, 185, 45, 255, 91, 95, 128}

	if !bytes.Equal(mint.Minter, minter) {
		return false
	}
	if !bytes.Equal(mint.MintAmount, amount) {
		return false
	}
	if !bytes.Equal(mint.MintTokens, tokens) {
		return false
	}
	return true
}

func checkRedeem(message *kafkautils.Message) bool {
	redeem := reflect.ValueOf(message.ProtoMsg).Interface().(*Redeem)
	redeemer := []byte{210, 95, 40, 189, 95, 239, 254, 5, 184, 50, 206, 247, 211, 47, 39, 199, 226, 34, 94, 17}
	redeemAmount := []byte{6, 73, 17, 26, 123, 208, 140, 161, 145}
	redeemTokens := []byte{131, 7, 100, 109, 133}

	if !bytes.Equal(redeem.Redeemer, redeemer) {
		return false
	}
	if !bytes.Equal(redeem.RedeemAmount, redeemAmount) {
		return false
	}
	if !bytes.Equal(redeem.RedeemTokens, redeemTokens) {
		return false
	}
	return true
}

func checkBorrow(message *kafkautils.Message) bool {
	borrow := reflect.ValueOf(message.ProtoMsg).Interface().(*Borrow)
	borrower := []byte{245, 93, 122, 47, 85, 59, 224, 190, 174, 220, 233, 3, 16, 58, 42, 19, 233, 181, 80, 140}
	borrowAmount := []byte{156, 63, 9, 125, 84, 143, 3, 112, 0, 0}
	accountBorrows := []byte{156, 63, 9, 125, 84, 143, 3, 112, 0, 0}
	totalBorrows := []byte{17, 21, 160, 249, 9, 78, 209, 122, 161, 32, 91}

	if !bytes.Equal(borrow.Borrower, borrower) {
		return false
	}
	if !bytes.Equal(borrow.BorrowAmount, borrowAmount) {
		return false
	}
	if !bytes.Equal(borrow.AccountBorrows, accountBorrows) {
		return false
	}
	if !bytes.Equal(borrow.TotalBorrows, totalBorrows) {
		return false
	}
	return true
}

func checkRepayBorrow(message *kafkautils.Message) bool {
	rb := reflect.ValueOf(message.ProtoMsg).Interface().(*RepayBorrow)
	payer := []byte{78, 57, 10, 0, 168, 57, 33, 55, 230, 163, 48, 249, 0, 163, 62, 148, 60, 62, 92, 10}
	borrower := []byte{78, 57, 10, 0, 168, 57, 33, 55, 230, 163, 48, 249, 0, 163, 62, 148, 60, 62, 92, 10}
	repayAmount := []byte{14, 241, 178, 220, 99, 119, 210, 203}
	var accountBorrows []byte
	totalBorrows := []byte{16, 89, 151, 132, 210, 158, 84, 29, 136, 191, 129}

	if !bytes.Equal(rb.Payer, payer) {
		return false
	}
	if !bytes.Equal(rb.Borrower, borrower) {
		return false
	}
	if !bytes.Equal(rb.RepayAmount, repayAmount) {
		return false
	}
	if !bytes.Equal(rb.AccountBorrows, accountBorrows) {
		return false
	}
	if !bytes.Equal(rb.TotalBorrows, totalBorrows) {
		return false
	}
	return true
}

func buildTopics(topicStr string) []ethcommon.Hash {
	t, _ := hexutil.Decode(topicStr)
	var b ethcommon.Hash
	copy(b[:], t)

	return []ethcommon.Hash{b}
}

func buildData(dataStr string) []byte {
	b, _ := hexutil.Decode(dataStr)
	return b
}
