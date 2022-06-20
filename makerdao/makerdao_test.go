package makerdao

import (
	"strings"
	"testing"

	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/makerdao/ETH"
	"github.com/nakji-network/connectors/makerdao/MCD_DAI"
	"github.com/nakji-network/connectors/makerdao/UNI"
	"github.com/nakji-network/connectors/makerdao/USDT"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockSubscription struct{}

var mockConnector *Connector

func (MockSubscription) Done() <-chan bool           { return nil }
func (MockSubscription) Err() <-chan error           { return nil }
func (MockSubscription) Headers() chan *types.Header { return nil }
func (MockSubscription) Logs() chan types.Log        { return nil }
func (MockSubscription) Resubscribe()                {}
func (MockSubscription) Unsubscribe()                {}
func (MockSubscription) GetBlockTime(vLog types.Log) (uint64, error) {
	cache := map[uint64]uint64{
		uint64(13145410): uint64(1629274175),
		uint64(13145411): uint64(1629274180),
	}
	return cache[vLog.BlockNumber], nil
}

func TestMain(*testing.M) {
	mockConnector = &Connector{
		Config: &Config{
			// ProtocolName:  "makerdao",
			ConnectorName: "makerdao_ethereum",
			NetworkName:   "ethereum",
		},
		sub:       new(MockSubscription),
		contracts: getMockContracts(),
	}

	registerTopics()
}

func TestParse(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		input *types.Log
		want  protoreflect.ProtoMessage
	}{
		{
			input: &types.Log{
				Address:     common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				BlockHash:   common.HexToHash("0x11a8fad69e2a6ceb2782045dfdf889217c1b893fb96bdda96d524aa1b32022af"),
				BlockNumber: 13145843,
				Data:        hexutil.MustDecode("0x00000000000000000000000011790264242190104"),
				Index:       157,
				TxIndex:     109,
				TxHash:      common.HexToHash("0xebffd1ca906353304166da64211430a5651edb564ddc7611a288cf5ab2254bb6"),
				Topics: []common.Hash{
					common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
					common.HexToHash("0x0000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d"),
				},
			},
			want: &ETH.Withdrawal{
				Ts: &timestamppb.Timestamp{
					Seconds: 1630580307,
				},
				Src: []byte("z%\rV0\xb4\xcfS\x979\xdf,]\xac\xb4\xc6Y\xf2H\x8d"),
				Wad: []byte("\x01\x11Z\x02@\xf2\xbeh"),
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				BlockHash:   common.HexToHash("0x11a8fad69e2a6ceb2782045dfdf889217c1b893fb96bdda96d524aa1b32022af"),
				BlockNumber: 13145843,
				Data:        hexutil.MustDecode("0x0000000000000000000000002241081542261622240"),
				Index:       130,
				TxIndex:     91,
				TxHash:      common.HexToHash("0xf2d239258a9a190e7ba1833275073f9d7018af300fbf968196a80a014e6d2c7e"),
				Topics: []common.Hash{
					common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
					common.HexToHash("0x0000000000000000000000001a2a1c938ce3ec39b6d47113c7955baa9dd454f2"),
				},
			},
			want: &ETH.Deposit{
				Ts: &timestamppb.Timestamp{
					Seconds: 1630580307,
				},
				Dst: []byte("\x1a*\x1c\x93\x8c\xe3\xec9\xb6\xd4q\x13Ǖ[\xaa\x9d\xd4T\xf2"),
				Wad: []byte("\x02\x18l\x9a\xe2\xa2\xe0\x00"),
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
				BlockHash:   common.HexToHash("0x11a8fad69e2a6ceb2782045dfdf889217c1b893fb96bdda96d524aa1b32022af"),
				BlockNumber: 13145843,
				Data:        hexutil.MustDecode("0x000000000000000000000041529125421212024814400"),
				Index:       47,
				TxIndex:     35,
				TxHash:      common.HexToHash("0xf4f60cf66e67aa31ec2d7ca032803e53b81d33d8f4dc69f45a3d1257f14002d3"),
				Topics: []common.Hash{
					common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					common.HexToHash("0x0000000000000000000000000cfeb7b8b2cf70e9e6fe768e43b8efbe640cc9ff"),
					common.HexToHash("0x0000000000000000000000003c9ff3cc55c82c82f4921083c1f32211d58225f5"),
				},
			},
			want: &MCD_DAI.Transfer{
				Ts: &timestamppb.Timestamp{
					Seconds: 1630580307,
				},
				Src: []byte("\x0c\xfe\xb7\xb8\xb2\xcfp\xe9\xe6\xfev\x8eC\xb8\xef\xbed\x0c\xc9\xff"),
				Dst: []byte("<\x9f\xf3\xccU\xc8,\x82\xf4\x92\x10\x83\xc1\xf3\"\x11Ղ%\xf5"),
				Wad: []byte("\x04\x98[\xfe\xd4x\xf8\x90\x00\x00"),
			},
		},
	} {
		res := mockConnector.parse(*testCase.input)
		if testCase.want != res {
			t.Error("Event log PROTOMSG parse failed.", "got:", res, "want:", testCase.want)
		}
	}
}

func getMockContracts() map[string]*Contract {
	uniABI, _ := abi.JSON(strings.NewReader(ABIs["UNI"]))
	ethABI, _ := abi.JSON(strings.NewReader(ABIs["ETH"]))
	usdtABI, _ := abi.JSON(strings.NewReader(ABIs["USDT"]))

	return map[string]*Contract{
		"0xdAC17F958D2ee523a2206206994597C13D831ec7": {
			ABI:  &uniABI,
			Name: "UNI",
			Type: "UNI",
		},
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2": {
			ABI:  &ethABI,
			Name: "ETH",
			Type: "ETH",
		},
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48": {
			ABI:  &usdtABI,
			Name: "USDT",
			Type: "USDT",
		},
	}
}

func registerTopics() {
	topicTypes := map[string]proto.Message{
		"nakji.makerdao.0_0_0.uni_transfer":  &UNI.Transfer{},
		"nakji.makerdao.0_0_0.eth_deposit":   &ETH.Deposit{},
		"nakji.makerdao.0_0_0.usdt_transfer": &USDT.Transfer{},
	}
	kafkautils.TopicTypeRegistry.Load(topicTypes)
}
