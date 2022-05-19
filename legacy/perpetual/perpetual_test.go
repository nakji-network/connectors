package perpetual

import (
	"strings"
	"testing"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/Amm"
	"blep.ai/data/connectors/source/perpetual/smart-contracts/L2/ClearingHouse"
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
		uint64(17645362): uint64(1629274175),
	}
	return cache[vLog.BlockNumber], nil
}

func TestMain(*testing.M) {
	mockConnector = &Connector{
		Config: &Config{
			Env:          "test",
			Author:       "nakji",
			ProtocolName: "perpetual",
			NetworkName:  "ethereum",
		},
		topicPrefix: ".fct.nakji.perpetual.",
		sub:         new(MockSubscription),
		contracts:   getMockContracts(),
	}

	registerTopics()
}

func TestParse(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		input *types.Log
		want  *kafkautils.Message
	}{
		{
			input: &types.Log{
				Address:     common.HexToAddress("0x187C938543f2BDE09Fe39034Fe3Ff797A3D35cA0"),
				BlockHash:   common.HexToHash("0xa04163bf1b759218bff8fbbe6e6c50ede6521317a1b2e4a77b07eddf7b327e85"),
				BlockNumber: 17645362,
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d8d726b7177a7fffff0000000000000000000000000000000000000000000000fc073cd3de858cc0c9"),
				Index:       35,
				TxIndex:     10,
				TxHash:      common.HexToHash("0x777eaa0836acea84bd0a76c8cc5e6dba35b3013420c91481d7c437affe93518f"),
				Topics: []common.Hash{
					common.HexToHash("0x0dd4066b1a6ce97fb670c3e4201e908c644193f38cbdaffd0229d7e26da3e533"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("perpetual", "0x187C938543f2BDE09Fe39034Fe3Ff797A3D35cA0"),
				Topic: kafkautils.MustParseTopic(".fct.nakji.perpetual.0_0_0.grtusdc_swapoutput", mockConnector.Env),
				ProtoMsg: &Amm.SwapOutput{
					Ts: &timestamppb.Timestamp{
						Seconds: 1629274175,
					},
					QuoteAssetAmount: []byte("\xd8\xd7&\xb7\x17z\x7f\xff\xff"),
					BaseAssetAmount:  []byte("\xfc\x07<\xd3ޅ\x8c\xc0\xc9"),
				},
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0x5d9593586b4B5edBd23E7Eba8d88FD8F09D83EBd"),
				BlockHash:   common.HexToHash("0xa04163bf1b759218bff8fbbe6e6c50ede6521317a1b2e4a77b07eddf7b327e85"),
				BlockNumber: 17645362,
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000002b5e3af16b188000000000000000000000000000000000000000000000000000d8d726b7177a800000fffffffffffffffffffffffffffffffffffffffffffffff6da48d95a2efadb6f0000000000000000000000000000000000000000000000003782dace9d900000fffffffffffffffffffffffffffffffffffffffffffffff6da48d95a2efadb6f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000148c33c7a9d9464e20000000000000000000000000000000000000000000000000000000000000000"),
				Index:       33,
				TxIndex:     9,
				TxHash:      common.HexToHash("0x0b5eed4531bdd58cd14553a4500d2b35e835d9c6c9f01fd3a39c846b5c7a4626"),
				Topics: []common.Hash{
					common.HexToHash("0x4c7b764f428c13bbea8cc8da90ebe6eef4dafeb27a4e3d9041d64208c47ca7c2"),
					common.HexToHash("0x0000000000000000000000007845ee7dbd6361689bafc6ebd287a7ee2a8a7436"),
					common.HexToHash("0x0000000000000000000000006de775aabeeede8efdb1a257198d56a3ac18c2fd"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("perpetual", "0x5d9593586b4B5edBd23E7Eba8d88FD8F09D83EBd"),
				Topic: kafkautils.MustParseTopic(".fct.perpetual.0_0_0.clearinghouse_positionchanged", mockConnector.Env),
				ProtoMsg: &ClearingHouse.PositionChanged{
					Ts: &timestamppb.Timestamp{
						Seconds: 1629274175,
					},
					Trader:           []byte("xE\xee}\xbdcah\x9b\xaf\xc6\xeb҇\xa7\xee*\x8at6"),
					Amm:              []byte("m\xe7u\xaa\xbe\xeeގ\xfd\xb1\xa2W\x19\x8dV\xa3\xac\x18\xc2\xfd"),
					Margin:           []byte("+^:\xf1k\x18\x80\x00\x00"),
					PositionNotional: []byte("\xd8\xd7&\xb7\x17z\x80\x00\x00"), ExchangedPositionSize: []byte("\t%\xb7&\xa5\xd1\x05$\x91"),
					Fee:               []byte("7\x82\xdaΝ\x90\x00\x00"),
					PositionSizeAfter: []byte("\t%\xb7&\xa5\xd1\x05$\x91"),
					SpotPrice:         []byte("\x01H\xc3<z\x9d\x94d\xe2"),
				},
			},
		},
		{
			input: &types.Log{
				Address:     common.HexToAddress("0x922F28072BaBe6EA0C0c25cCD367Fda0748a5EC7"),
				BlockHash:   common.HexToHash("0xa04163bf1b759218bff8fbbe6e6c50ede6521317a1b2e4a77b07eddf7b327e85"),
				BlockNumber: 17645362,
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000004e17c98a12b3f72aa5bef0000000000000000000000000000000000000000000838265642b3601faa5c3d00000000000000000000000000000000000000000000000000000000611cc03f"),
				Index:       5,
				TxIndex:     1,
				TxHash:      common.HexToHash("0x736b4eb3237705b7c84fd3b3d673bd5dc333cf28c412d9e73af1b3fc8124f05b"),
				Topics: []common.Hash{
					common.HexToHash("0x3a3348362552c3897fd1f06a3233519ebd8bd76ad6e99a418a9741155fe90515"),
				},
			},
			want: &kafkautils.Message{
				Key:   kafkautils.NewKey("perpetual", "0x922F28072BaBe6EA0C0c25cCD367Fda0748a5EC7"),
				Topic: kafkautils.MustParseTopic(".fct.perpetual.0_0_0.renusdc_reservesnapshotted", mockConnector.Env),
				ProtoMsg: &Amm.ReserveSnapshotted{
					Ts: &timestamppb.Timestamp{
						Seconds: 1629274175,
					},
					QuoteAssetReserve: []byte("\x04\xe1|\x98\xa1+?r\xaa[\xef"),
					BaseAssetReserve:  []byte("\x088&VB\xb3`\x1f\xaa\\="),
					Timestamp:         []byte("a\x1c\xc0?"),
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
	ammABI, _ := abi.JSON(strings.NewReader(ABIs["Amm"]))
	clearingHouseABI, _ := abi.JSON(strings.NewReader(ABIs["ClearingHouse"]))

	return map[string]*Contract{
		"0x187C938543f2BDE09Fe39034Fe3Ff797A3D35cA0": {
			ABI:  &ammABI,
			Name: "Amm",
		},
		"0x5d9593586b4B5edBd23E7Eba8d88FD8F09D83EBd": {
			ABI:  &clearingHouseABI,
			Name: "ClearingHouse",
		},
	}
}

func registerTopics() {
	topicTypes := map[string]proto.Message{
		"nakji.perpetual.0_0_0.grtusdc_swapoutput":            &Amm.SwapOutput{},
		"nakji.perpetual.0_0_0.clearinghouse_positionchanged": &ClearingHouse.PositionChanged{},
		"nakji.perpetual.0_0_0.renusdc_reservesnapshotted":    &Amm.ReserveSnapshotted{},
	}
	kafkautils.TopicTypeRegistry.Load(topicTypes)
}
