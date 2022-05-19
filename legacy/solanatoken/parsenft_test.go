package solanatoken

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "blep.ai/data/chain/solana/testing"
	"google.golang.org/protobuf/types/known/timestamppb"

	bin "github.com/gagliardetto/binary"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	mint        solana.PublicKey
	invalidMint solana.PublicKey
)

func TestParseNFT(t *testing.T) {
	payer := NewPubKey(t)
	mint = NewPubKey(t)
	invalidMint = NewPubKey(t)

	notNFTmintTo1 := BuildTokenMintTo(12, mint, payer, payer, []solana.PublicKey{})
	invalidData1, _ := notNFTmintTo1.Data()

	notNFTmintTo2 := BuildTokenMintTo(1, invalidMint, payer, payer, []solana.PublicKey{})
	invalidData2, _ := notNFTmintTo2.Data()

	mintTo := BuildTokenMintTo(1, mint, payer, payer, []solana.PublicKey{})
	mintToData, _ := mintTo.Data()

	tests := []struct {
		name string
		args []testInstruction
		want *NFT
	}{
		{
			name: "MintTo does not exist",
			args: []testInstruction{
				&testTxInstruction{
					accounts: []*solana.AccountMeta{
						{PublicKey: payer, IsSigner: true, IsWritable: false},
					},
					programID: solana.TokenMetadataProgramID,
				},
			},
			want: nil,
		},
		{
			name: "MintTo exists but amount minted greater than 1",
			args: []testInstruction{
				&testTxInstruction{
					accounts: []*solana.AccountMeta{
						{PublicKey: payer, IsSigner: true, IsWritable: false},
					},
					programID: solana.TokenMetadataProgramID,
				},
				&testTxInstruction{
					accounts:          notNFTmintTo1.Accounts(),
					programID:         notNFTmintTo1.ProgramID(),
					data:              invalidData1,
					innerInstructions: []solana.Instruction{},
				},
			},
			want: nil,
		},
		{
			name: "MintTo exists and amount minted greater than 1 but not an NFT",
			args: []testInstruction{
				&testTxInstruction{
					accounts: []*solana.AccountMeta{
						{PublicKey: payer, IsSigner: true, IsWritable: false},
					},
					programID: solana.TokenMetadataProgramID,
				},
				&testTxInstruction{
					accounts:          notNFTmintTo2.Accounts(),
					programID:         notNFTmintTo2.ProgramID(),
					data:              invalidData2,
					innerInstructions: []solana.Instruction{},
				},
			},
			want: nil,
		},
		{
			name: "MintTo exists for a valid NFT",
			args: []testInstruction{
				&testTxInstruction{
					accounts: []*solana.AccountMeta{
						{PublicKey: payer, IsSigner: true, IsWritable: false},
					},
					programID: solana.TokenMetadataProgramID,
				},
				&testTxInstruction{
					accounts:          mintTo.Accounts(),
					programID:         mintTo.ProgramID(),
					data:              mintToData,
					innerInstructions: []solana.Instruction{},
				},
			},
			want: &NFT{
				Address:             mint.String(),
				Ts:                  &timestamppb.Timestamp{Seconds: 1000},
				MintAuthority:       payer.String(),
				FreezeAuthority:     "",
				UpdateAuthority:     payer.String(),
				PrimarySaleHappened: true,
				Royalty:             1000,
				Creators:            []*Creator{},
				Name:                "Foo",
				Symbol:              "FOO",
				Uri:                 "www.foo.com",
				Collection:          "",
			},
		},
	}

	c, closer := newParseNFTConnector(payer)
	defer closer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			txResult, tx := NewTxResult(tt.args)
			got := c.parseNFT(txResult, tx)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNFT() got = %+v, \nwant = %+v", got, tt.want)
			}
		})
	}
}

func newParseNFTConnector(authority solana.PublicKey) (*SolanaTokenConnector, func()) {
	server, closer := mockNFTJSONRPC(authority)
	client := rpc.New(server.URL)
	connector := testConnector(client)
	return connector, closer
}

type Request struct {
	Method string
	Params []interface{}
}

func mockNFTJSONRPC(authority solana.PublicKey) (*MockJSONRPCServer, func()) {
	metaPubKey, _, _ := solana.FindTokenMetadataAddress(mint)

	mock := &MockJSONRPCServer{
		Server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			var r Request
			json.NewDecoder(req.Body).Decode(&r)

			var responseBody json.RawMessage
			pubKey := r.Params[0].(string)

			if r.Method != "getAccountInfo" {
				return
			}

			switch pubKey {
			case invalidMint.String():
				tokenMint := token.Mint{
					MintAuthority: &authority,
					Supply:        1,
					Decimals:      9,
					IsInitialized: true,
				}
				responseBody = wrapAccountResponse(tokenMint, bin.NewBorshEncoder)
			case mint.String():
				tokenMint := token.Mint{
					MintAuthority: &authority,
					Supply:        1,
					Decimals:      0,
					IsInitialized: true,
				}
				responseBody = wrapAccountResponse(tokenMint, bin.NewBorshEncoder)
			case metaPubKey.String():
				metadata := token_metadata.Metadata{
					UpdateAuthority: authority,
					Mint:            mint,
					Data: token_metadata.Data{
						Name:                 "Foo",
						Symbol:               "FOO",
						Uri:                  "www.foo.com",
						SellerFeeBasisPoints: 1000,
						Creators:             &[]token_metadata.Creator{},
					},
					PrimarySaleHappened: true,
					IsMutable:           true,
				}
				responseBody = wrapAccountResponse(metadata, bin.NewBorshEncoder)
			}

			rw.Write(responseBody)
		})),
	}

	return mock, func() { mock.Close() }
}
