package solanatoken

import (
	"reflect"
	"testing"

	. "blep.ai/data/chain/solana/testing"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func TestParseTrade(t *testing.T) {
	pubKey := NewPubKey(t)
	mint := solana.MustPublicKeyFromBase58("AH5EmaU1NSkbtR3WW6diCKkQx1cGuJFs8L28GKHLyGmj")
	initializeAccount := BuildInitializeAccountInstruction(pubKey, mint, pubKey)
	createAssocAccountInstructions := []solana.Instruction{initializeAccount}
	buyer := NewPubKey(t)
	seller := NewPubKey(t)

	exchangeArtBuyInstructions := []solana.Instruction{
		// Sol transfer for Transaction fee paid to NFT Marketplace
		BuildSolTransferInstruction(2000, pubKey, pubKey),
		// Sol transfer for 2-3 Creators that receive royalties
		BuildSolTransferInstruction(50000, pubKey, pubKey),
		BuildSolTransferInstruction(50000, pubKey, pubKey),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(100000, buyer, seller),
		// 2nd to last instruction is TokenTransfer
		BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
		// last instruction is TokenClose
		BuildTokenCloseInstruction(pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
	}
	exchangeArtInsts := []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: mint},
				{PublicKey: solana.TokenProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
			innerInstructions: createAssocAccountInstructions,
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(ExchangeArtProgramId),
			innerInstructions: exchangeArtBuyInstructions,
		},
	}

	// https://solscan.io/tx/25PVi5CwHDqJmUd8H3kUxKmqsNq4DxWXyarqYciBFwB21LAxTkUwCPXN6jiinsfk8HBnAa3oZwLtnquH1SoLWELn
	magicEdenV2BuyInsts := []solana.Instruction{
		// Sol transfer for 1-3 Creators that receive royalties
		BuildSolTransferInstruction(0.065*1e9, pubKey, pubKey), // This first transfer is usually 0
		// Sol transfer for 2% transaction fee paid to MagicEden
		BuildSolTransferInstruction(0.013*1e9, pubKey, pubKey),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(0.572*1e9, buyer, seller),
		BuildCreateAssociatedAccount(buyer, pubKey, pubKey),
		// Fixed fee that seems to be returned to buyer
		BuildSolTransferInstruction(0.00203928*1e9, pubKey, pubKey),
		BuildSystemAssignInstruction(pubKey, pubKey),
		BuildInitializeAccountInstruction(pubKey, pubKey, pubKey),
		BuildTokenTransferInstruction(1, seller, buyer, seller, []solana.PublicKey{seller}),
		BuildTokenCloseInstruction(pubKey, buyer, pubKey, []solana.PublicKey{seller}),
	}

	magicEdenEscrowPubKey := solana.MustPublicKeyFromBase58(MagicEdenEscrowAccount)
	magicEdenBuyInsts := []solana.Instruction{
		// Sol transfer for 2% transaction fee paid to MagicEden
		BuildSolTransferInstruction(1.5*1e9, pubKey, pubKey),
		// Sol transfer for 2-3 Creators that receive royalties
		BuildSolTransferInstruction(0, pubKey, pubKey), // This first transfer is usually 0
		BuildSolTransferInstruction(3.75*1e9, pubKey, pubKey),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(uint64(69.75*1e9), buyer, seller),
		// Last instruction is SetAuthority instead of the usual Token Transfer
		// subject (2nd arg) is the token account that has the NFT. authority (3rd arg) is Magic Eden's
		// escrow account (https://solscan.io/account/GUfCR9mK6azb9vcpsxgXyj7XRPAKJd4KMHTTVvtncGgp)
		BuildTokenSetAuthority(buyer, pubKey, magicEdenEscrowPubKey, []solana.PublicKey{buyer}),
	}

	// These instructions mostly map to https://solscan.io/tx/Y3oaunErxoBdqwzDbwGNHNhjJ9A8WoEJvXb3XM4r3F9NTCvnnkWhRzfo71WCJtM2zbkXrjWrkqe2SiaMTso3ZtJ
	solanartBuyInsts := []solana.Instruction{
		// Sol transfer for 3 Creators that receive royalties
		BuildSolTransferInstruction(0.02772*1e9, pubKey, pubKey),
		BuildSolTransferInstruction(0.02772*1e9, pubKey, pubKey),
		BuildSolTransferInstruction(0.02856*1e9, pubKey, pubKey),
		// Sol transfer for 3% transaction fee paid to Solanart
		BuildSolTransferInstruction(0.0336*1e9, pubKey, solana.MustPublicKeyFromBase58(SolanartFeeAddress)),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(uint64(1.0024*1e9), buyer, seller),
		// 2nd to last instruction is TokenTransfer. Transfers NFT from seller's TokenAccount to buyer's TokenAccount
		BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
		// Last instruction is CloseAccount which closes the TokenAccount the NFT was transferred from
		BuildTokenCloseInstruction(pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
	}

	solanartInsts := []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: mint},
				{PublicKey: solana.TokenProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
			innerInstructions: createAssocAccountInstructions,
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58(SolanartFeeAddress)},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(SolanartProgramId),
			innerInstructions: solanartBuyInsts,
		},
	}

	// These instructions mostly map to https://solscan.io/tx/484ATxiDpTvzB9XLekCnYGDtAcKW5kb6Zy4civ3zYmUMffqaSnuBas6ubwumnp8rsJ2Gb4QXDBu3S6u1Tc6v6s5Y
	solseaBuyInsts := []solana.Instruction{
		// Sol transfer for 3% transaction fee paid to Solsea
		BuildSolTransferInstruction(0.006*1e9, pubKey, solana.MustPublicKeyFromBase58(SolseaFeeAddress)),
		// Sol transfer for 2 Creators that receive royalties
		BuildSolTransferInstruction(0.0291*1e9, pubKey, pubKey),
		BuildSolTransferInstruction(0, pubKey, pubKey),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(0.1649*1e9, buyer, seller),
		// Last instruction is TokenTransfer. Transfers NFT from seller's TokenAccount to buyer's TokenAccount
		BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
	}

	// These instructions mostly map to https://solscan.io/tx/tPyKk6ucj8Mca7r9mdUfzcz3D9rLcc4FYivS4nuRfEpkn7Q4PywWgsHkECYizNzYFX46hNwbFQgJRynE26K4HL3
	digitalEyesBuyInsts := []solana.Instruction{
		// Sol transfer for 2.5% transaction fee paid to DigitalEyes
		BuildSolTransferInstruction(0.00625*1e9, pubKey, solana.MustPublicKeyFromBase58(DigitalEyesFeeAddress)),
		// Sol transfer for 3 Creators that receive royalties
		BuildSolTransferInstruction(0, pubKey, pubKey),
		BuildSolTransferInstruction(0.0125*1e9, pubKey, pubKey),
		// Sol transfer with Buyer and Seller info
		BuildSolTransferInstruction(uint64(0.23125*1e9), buyer, seller),
		// 2nd to last instruction is TokenTransfer. Transfers NFT from seller's TokenAccount to buyer's TokenAccount
		BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
		// Last instruction is CloseAccount which closes the TokenAccount the NFT was transferred from
		BuildTokenCloseInstruction(pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
	}

	digitalEyesInsts := []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: mint},
				{PublicKey: solana.TokenProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
			innerInstructions: createAssocAccountInstructions,
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58(DigitalEyesFeeAddress)},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(DigitalEyesProgramId),
			innerInstructions: digitalEyesBuyInsts,
		},
	}

	tests := []struct {
		name string
		args []testInstruction
		want *Trade
	}{
		{
			name: "MagicEdenV2 - invalid since 2nd to last instruction is not a Token Transfer",
			args: magicEdenV2Instructions(buyer, seller, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildSolTransferInstruction(0.065*1e9, pubKey, pubKey),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			}),
			want: nil,
		},
		{
			name: "MagicEdenV2 - invalid since buy instructions are less than 8",
			args: magicEdenV2Instructions(buyer, seller, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			}),
			want: nil,
		},
		{
			name: "MagicEdenV2 - invalid since instructions are neither 2 nor 3",
			args: []testInstruction{
				&testTxInstruction{
					accounts: []*solana.AccountMeta{
						{PublicKey: buyer, IsSigner: true, IsWritable: false},
					},
					data:              []byte{0xaa, 0xbb},
					programID:         solana.MustPublicKeyFromBase58(MagicEdenV2ProgramId),
					innerInstructions: []solana.Instruction{},
				},
			},
			want: nil,
		},
		{
			name: "MagicEdenV2 - parses valid trade",
			args: magicEdenV2Instructions(buyer, seller, pubKey, magicEdenV2BuyInsts),
			want: &Trade{
				NftAddress:  pubKey.String(),
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "magicedenv2",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       0.65 * 1e9,
			},
		},
		{
			name: "MagicEden - invalid since less than 4 buy instructions",
			args: magicEdenInstructions(buyer, seller, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			}),
			want: nil,
		},
		{
			name: "MagicEden - invalid since last buy instruction is not token SetAuthority",
			args: magicEdenInstructions(buyer, seller, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			}),
			want: nil,
		},
		{
			name: "MagicEden - invalid since current authority in token.SetAuthority must be Magic Eden's Escrow Account",
			args: magicEdenInstructions(buyer, seller, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				BuildTokenSetAuthority(buyer, pubKey, pubKey, []solana.PublicKey{buyer}),
			}),
			want: nil,
		},
		{
			name: "MagicEden - parses valid trade",
			args: magicEdenInstructions(buyer, seller, pubKey, magicEdenBuyInsts),
			want: &Trade{
				NftAddress:  "BGKNQe6pBFfaDhjMmcj61VeST3hVg1iQjE1tgMqNrQDz",
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "magiceden",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       75 * 1e9,
			},
		},
		{
			name: "Solsea - invalid since less than 4 buy instructions",
			args: solseaInstructions(buyer, seller, mint, pubKey, []solana.Instruction{
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			}),
			want: nil,
		},
		{
			name: "Solsea - invalid since last buy instruction is not token transfer",
			args: solseaInstructions(buyer, seller, mint, pubKey, []solana.Instruction{
				BuildSolTransferInstruction(2000, pubKey, pubKey),
			}),
			want: nil,
		},
		{
			name: "Solsea - invalid since first buy instruction is not SOL transfer to Solsea's Fee address",
			args: solseaInstructions(buyer, seller, mint, pubKey, []solana.Instruction{
				BuildSolTransferInstruction(2000, pubKey, pubKey),
			}),
			want: nil,
		},
		{
			name: "Solsea - parses valid trade",
			args: solseaInstructions(buyer, seller, mint, pubKey, solseaBuyInsts),
			want: &Trade{
				NftAddress:  mint.String(),
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "solsea",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       0.2 * 1e9,
			},
		},
		{
			name: "Solanart - invalid with no token close inst",
			args: noTokenClose(buyer, seller, mint, pubKey, SolanartProgramId),
			want: nil,
		},
		{
			name: "Solanart - invalid since less than 4 buy instructions",
			args: lessThan4BuyInsts(buyer, seller, mint, pubKey, SolanartProgramId),
			want: nil,
		},
		{
			name: "Solanart - invalid since 2nd to last buy instruction is not token transfer",
			args: noTokenTransfer(buyer, seller, mint, pubKey, SolanartProgramId),
			want: nil,
		},
		{
			name: "Solanart - invalid since no initializeAccount instruction",
			args: noInitializeAcct(buyer, seller, mint, pubKey, SolanartProgramId, solanartBuyInsts),
			want: nil,
		},
		{
			name: "Solanart - parses valid trade",
			args: solanartInsts,
			want: &Trade{
				NftAddress:  mint.String(),
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "solanart",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       1.12 * 1e9,
			},
		},
		{
			name: "DigitalEyes - invalid with no token close inst",
			args: noTokenClose(buyer, seller, mint, pubKey, DigitalEyesProgramId),
			want: nil,
		},
		{
			name: "DigitalEyes - invalid since less than 4 buy instructions",
			args: lessThan4BuyInsts(buyer, seller, mint, pubKey, DigitalEyesProgramId),
			want: nil,
		},
		{
			name: "DigitalEyes - invalid since 2nd to last buy instruction is not token transfer",
			args: noTokenTransfer(buyer, seller, mint, pubKey, DigitalEyesProgramId),
			want: nil,
		},
		{
			name: "DigitalEyes - invalid since no initializeAccount instruction",
			args: noInitializeAcct(buyer, seller, mint, pubKey, DigitalEyesProgramId, digitalEyesBuyInsts),
			want: nil,
		},
		{
			name: "DigitalEyes - parses valid trade",
			args: digitalEyesInsts,
			want: &Trade{
				NftAddress:  mint.String(),
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "digitaleyes",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       0.25 * 1e9,
			},
		},
		{
			name: "ExchangeArt - invalid with no token close inst",
			args: noTokenClose(buyer, seller, mint, pubKey, ExchangeArtProgramId),
			want: nil,
		},
		{
			name: "ExchangeArt - invalid since less than 4 buy instructions",
			args: lessThan4BuyInsts(buyer, seller, mint, pubKey, ExchangeArtProgramId),
			want: nil,
		},
		{
			name: "ExchangeArt - invalid since 2nd to last buy instruction is not token transfer",
			args: noTokenTransfer(buyer, seller, mint, pubKey, ExchangeArtProgramId),
			want: nil,
		},
		{
			name: "ExchangeArt - invalid since no initializeAccount instruction",
			args: noInitializeAcct(buyer, seller, mint, pubKey, ExchangeArtProgramId, exchangeArtBuyInstructions),
			want: nil,
		},
		{
			name: "ExchangeArt - parses valid trade",
			args: exchangeArtInsts,
			want: &Trade{
				NftAddress:  mint.String(),
				Ts:          timestamppb.New(testTs.Time()),
				SoldOn:      "exchangeart",
				Transaction: testSignature.String(),
				Buyer:       buyer.String(),
				Seller:      seller.String(),
				Price:       202000,
			},
		},
	}

	c, closer := newParseTradeConnector(t)
	defer closer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			txResult, tx := NewTxResult(tt.args)
			got := c.parseTrade(txResult, tx)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTrade() got = \n%v\n, want = \n%v\n", got, tt.want)
			}
		})
	}
}

func noTokenClose(buyer, seller, mint, pubKey solana.PublicKey, nftMarketProgramId string) []testInstruction {
	return []testInstruction{
		validSolanartFirstInstruction(buyer, mint, pubKey),
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:      []byte{0xaa, 0xbb},
			programID: solana.MustPublicKeyFromBase58(nftMarketProgramId),
			innerInstructions: []solana.Instruction{
				// Sol transfer for Transaction fee paid to NFT Marketplace
				BuildSolTransferInstruction(2000, pubKey, pubKey),
				// Sol transfer for 2-3 Creators that receive royalties
				BuildSolTransferInstruction(50000, pubKey, pubKey),
				BuildSolTransferInstruction(50000, pubKey, pubKey),
				// Sol transfer with Buyer and Seller info
				BuildSolTransferInstruction(100000, buyer, seller),
				// 2nd to last instruction is TokenTransfer
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				// invalid because no TokenClose as last instruction
			},
		},
	}
}

func lessThan4BuyInsts(buyer, seller, mint, pubKey solana.PublicKey, nftMarketProgramId string) []testInstruction {
	return []testInstruction{
		validSolanartFirstInstruction(buyer, mint, pubKey),
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:      []byte{0xaa, 0xbb},
			programID: solana.MustPublicKeyFromBase58(nftMarketProgramId),
			innerInstructions: []solana.Instruction{
				// Invalid because total inner instructions here are less than 4
				// Sol transfer for Transaction fee paid to NFT Marketplace
				BuildSolTransferInstruction(2000, pubKey, pubKey),
				// Sol transfer with Buyer and Seller info
				BuildSolTransferInstruction(100000, buyer, seller),
				// 2nd to last instruction is TokenTransfer
				BuildTokenTransferInstruction(1, pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
				// invalid because no TokenClose as last instruction
			},
		},
	}
}

func noTokenTransfer(buyer, seller, mint, pubKey solana.PublicKey, nftMarketProgramId string) []testInstruction {
	return []testInstruction{
		validSolanartFirstInstruction(buyer, mint, pubKey),
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:      []byte{0xaa, 0xbb},
			programID: solana.MustPublicKeyFromBase58(nftMarketProgramId),
			innerInstructions: []solana.Instruction{
				// Invalid because total inner instructions here are less than 4
				// Sol transfer for Transaction fee paid to NFT Marketplace
				BuildSolTransferInstruction(2000, pubKey, pubKey),
				// Sol transfer with Buyer and Seller info
				BuildSolTransferInstruction(100000, buyer, seller),
				// 2nd to last instruction is supposed to be TokenTransfer
				BuildSolTransferInstruction(100000, buyer, seller),
				// invalid because no TokenClose as last instruction
				BuildTokenCloseInstruction(pubKey, pubKey, pubKey, []solana.PublicKey{pubKey}),
			},
		},
	}
}

func noInitializeAcct(
	buyer, seller, mint, pubKey solana.PublicKey,
	nftMarketProgramId string,
	validBuyInstructions []solana.Instruction,
) []testInstruction {
	return []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: mint},
				{PublicKey: solana.TokenProgramID},
			},
			data:      []byte{0xaa, 0xbb},
			programID: solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
			// invalid because last inner instruction here must be a Token.InitializeAccount inst
			innerInstructions: []solana.Instruction{BuildTokenCloseInstruction(pubKey, pubKey, pubKey, []solana.PublicKey{pubKey})},
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(nftMarketProgramId),
			innerInstructions: validBuyInstructions,
		},
	}
}

func validSolanartFirstInstruction(buyer, mint, pubKey solana.PublicKey) testInstruction {
	initializeAccount := BuildInitializeAccountInstruction(pubKey, mint, pubKey)
	createAssocAccountInstructions := []solana.Instruction{initializeAccount}

	return &testTxInstruction{
		accounts: []*solana.AccountMeta{
			{PublicKey: buyer, IsSigner: true, IsWritable: false},
			{PublicKey: mint},
			{PublicKey: solana.TokenProgramID},
		},
		data:              []byte{0xaa, 0xbb},
		programID:         solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
		innerInstructions: createAssocAccountInstructions,
	}
}

func solseaInstructions(buyer, seller, mint, pubKey solana.PublicKey, buyInstructions []solana.Instruction) []testInstruction {
	initializeAccount := BuildInitializeAccountInstruction(pubKey, mint, pubKey)
	data, _ := initializeAccount.Data()

	return []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{{PublicKey: buyer, IsSigner: true, IsWritable: false}},
		},
		&testTxInstruction{
			accounts:          initializeAccount.Accounts(),
			data:              data,
			programID:         initializeAccount.ProgramID(),
			innerInstructions: []solana.Instruction{},
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
				{PublicKey: solana.MustPublicKeyFromBase58(SolseaFeeAddress)},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(SolseaProgramId),
			innerInstructions: buyInstructions,
		},
	}
}

func magicEdenV2Instructions(buyer, seller, pubKey solana.PublicKey, buyInstructions []solana.Instruction) []testInstruction {
	return []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(MagicEdenV2ProgramId),
			innerInstructions: []solana.Instruction{},
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(MagicEdenV2ProgramId),
			innerInstructions: []solana.Instruction{},
		},
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: pubKey},
				{PublicKey: pubKey},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(MagicEdenV2ProgramId),
			innerInstructions: buyInstructions,
		},
	}
}

func magicEdenInstructions(buyer, seller, pubKey solana.PublicKey, buyInstructions []solana.Instruction) []testInstruction {
	magicEdenEscrowPubKey := solana.MustPublicKeyFromBase58(MagicEdenEscrowAccount)
	return []testInstruction{
		&testTxInstruction{
			accounts: []*solana.AccountMeta{
				{PublicKey: buyer, IsSigner: true, IsWritable: false},
				{PublicKey: seller},
				{PublicKey: pubKey},
				{PublicKey: solana.MustPublicKeyFromBase58("2nzX2GrJdUu9BzHqgecG22F5yikmmkarVpLSic6vekyo")},
				{PublicKey: solana.TokenProgramID},
				{PublicKey: solana.SystemProgramID},
				{PublicKey: magicEdenEscrowPubKey},
			},
			data:              []byte{0xaa, 0xbb},
			programID:         solana.MustPublicKeyFromBase58(MagicEdenProgramId),
			innerInstructions: buyInstructions,
		},
	}
}

func newParseTradeConnector(t *testing.T) (*SolanaTokenConnector, func()) {
	tokenAccount := token.Account{
		Mint: solana.MustPublicKeyFromBase58("BGKNQe6pBFfaDhjMmcj61VeST3hVg1iQjE1tgMqNrQDz"),
	}
	resp := wrapAccountResponse(tokenAccount, bin.NewBinEncoder)
	server, closer := MockJSONRPC(t, resp)
	client := rpc.New(server.URL)
	connector := testConnector(client)
	return connector, closer
}
