package starknet

import (
	"math/big"
	"strings"
	"testing"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/connectors/source/ethstarknet/smart-contracts/Starknet"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockSubscription struct{}

var mockConnector *Connector

func (MockSubscription) ClientPool() *ethclient.ClientPool { return nil }
func (MockSubscription) Done() <-chan bool                 { return nil }
func (MockSubscription) Err() <-chan error                 { return nil }
func (MockSubscription) Headers() chan *types.Header       { return nil }
func (MockSubscription) Logs() chan types.Log              { return nil }
func (MockSubscription) Resubscribe()                      {}
func (MockSubscription) Unsubscribe()                      {}
func (MockSubscription) GetBlockTime(vLog types.Log) (uint64, error) {
	cache := map[uint64]uint64{
		uint64(13945205): uint64(1641382544),
		uint64(13945198): uint64(1641382464),
		uint64(13945197): uint64(1641382439),
	}
	return cache[vLog.BlockNumber], nil
}

func TestMain(*testing.M) {
	mockConnector = &Connector{
		Config: &Config{
			Env:          "test",
			Author:       "nakji",
			ProtocolName: "starknet",
			NetworkName:  "ethereum",
			Version:      "0_0_0",
		},
		topicPrefix: ".fct.nakji.starknet.",
		sub:         new(MockSubscription),
		contracts:   getMockContracts(),
	}

	registerTopics(mockConnector.Version)
}

func TestParse(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		input *types.Log
		want  *kafkautils.Message
	}{
		{
			input: &types.Log{
				Address:     common.HexToAddress("0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				BlockHash:   common.HexToHash("0xf504dcc1f36fbf368ac4b365f0b2bad04b4f93f405e6e98a3d8c593cdcd219f7"),
				BlockNumber: 13945205,
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000022"),
				Index:       278,
				TxIndex:     197,
				TxHash:      common.HexToHash("0xaa4d3eab0bd09837c717fa634c844614b2a53a48ac7ee29c0f29b2d74455973c"),
				Topics: []common.Hash{
					common.HexToHash("0x4264ac208b5fde633ccdd42e0f12c3d6d443a4f3779bbf886925b94665b63a22"),
					common.HexToHash("0x03b9ae7e554908cac546ec846951570566cc97c95e0387f59da09dbb91e9abf6"),
					common.HexToHash("0x000000000000000000000000adbad9db9ecd54c7dc9bbb1b65ab658539ef2ff8"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("starknet", "0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				Topic: kafkautils.MustParseTopic(".fct.nakji.starknet.0_0_0.starknet_logmessagetol1", mockConnector.Env),
				ProtoMsg: &Starknet.LogMessageToL1{
					Ts: &timestamppb.Timestamp{
						Seconds: 1641382544,
					},
					Fromaddress: []byte{7, 212, 195, 104, 217, 162, 61, 33, 177, 99, 254, 213, 114, 2, 23, 199, 211, 24, 204, 216, 196, 100, 224, 154, 91, 93, 179, 28, 202, 52, 174, 92},
					Toaddress:   common.HexToAddress("0xaDBAd9dB9ECD54c7dc9bBB1B65AB658539ef2fF8").Bytes(),
					Payload:     []byte{12, 34},
				},
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				BlockHash:   common.HexToHash("0x7cba733b5af64c3928c967fd29121388316c45a137b6ed2175a56c019dee8450"),
				BlockNumber: 13945198,
				Data:        hexutil.MustDecode("0x025f30c08ac40231f495a63894f14b40c02db2530b5cdd0dc71e826e3dcf4bee0000000000000000000000000000000000000000000000000000000000000249"),
				Index:       351,
				TxIndex:     231,
				TxHash:      common.HexToHash("0x9b9b9fcb7f89752ef4312b26b11f9c0a22775a32398b0de1be00af2516d7a048"),
				Topics: []common.Hash{
					common.HexToHash("0xe8012213bb931d3efa0a954cfb0d7b75f2a5e2358ba5f7d3edfb0154f6e7a568"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("starknet", "0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				Topic: kafkautils.MustParseTopic(".fct.nakji.starknet.0_0_0.starknet_logstateupdate", mockConnector.Env),
				ProtoMsg: &Starknet.LogStateUpdate{
					Ts: &timestamppb.Timestamp{
						Seconds: 1641382464,
					},
					GlobalRoot:  []byte{2, 95, 48, 192, 138, 196, 2, 49, 244, 149, 166, 56, 148, 241, 75, 64, 192, 45, 178, 83, 11, 92, 221, 13, 199, 30, 130, 110, 61, 207, 75, 238},
					BlockNumber: big.NewInt(13945198).Bytes(),
				},
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				BlockHash:   common.HexToHash("0xfe2db8a09acc1fb02588d8b874f04ffe4ebf59f7927378a28f5afade56962f24"),
				BlockNumber: 13945197,
				Data:        hexutil.MustDecode("0xd32ce06a9f6e96cbbce81e0137bf165daff445664966eb72eafc0a31a4dfa3cb"),
				Index:       257,
				TxIndex:     186,
				TxHash:      common.HexToHash("0x0b1f984d35a212a1051e9c56dae4b7eee8cefd6d7f987f3038741cfa4da23fdb"),
				Topics: []common.Hash{
					common.HexToHash("0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("starknet", "0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4"),
				Topic: kafkautils.MustParseTopic(".fct.nakji.starknet.0_0_0.starknet_logstatetransitionfact", mockConnector.Env),
				ProtoMsg: &Starknet.LogStateTransitionFact{
					Ts: &timestamppb.Timestamp{
						Seconds: 1641382439,
					},
					StateTransitionFact: common.Hex2Bytes("0xd32ce06a9f6e96cbbce81e0137bf165daff445664966eb72eafc0a31a4dfa3cb"),
				},
			},
		},
	} {
		res := mockConnector.parse(*testCase.input)
		if testCase.want.Key.String() != res.Key.String() {
			t.Error("Event log KEY parse failed.", "got:", res.Key, "want:", testCase.want.Key)
		}

		if testCase.want.Topic.String() != res.Topic.String() {
			t.Error("Event log TOPIC parse failed.", "got:", res.Key, "want:", testCase.want.Key)
		}

		if testCase.want.ProtoMsg.ProtoReflect().Descriptor().Syntax().String() != res.ProtoMsg.ProtoReflect().Descriptor().Syntax().String() {
			t.Error("Event log PROTOMSG parse failed.", "got:", res.Key, "want:", testCase.want.Key)
		}
	}
}

func getMockContracts() map[string]*Contract {
	starknetABI, _ := abi.JSON(strings.NewReader(ABIs["Starknet"]))

	return map[string]*Contract{
		"0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4": {
			ABI:  &starknetABI,
			Name: "Starknet",
		},
	}
}

func registerTopics(version string) {
	topicTypes := map[string]proto.Message{
		"nakji.starknet.0_0_0.starknet_logmessagetol1":         &Starknet.LogMessageToL1{},
		"nakji.starknet.0_0_0.starknet_logstateupdate":         &Starknet.LogStateUpdate{},
		"nakji.starknet.0_0_0.starknet_logstatetransitionfact": &Starknet.LogStateTransitionFact{},
	}
	kafkautils.TopicTypeRegistry.Load(topicTypes)
}
