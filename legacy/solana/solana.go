package solana

import (
	"context"
	"fmt"
	"time"

	"blep.ai/data/connectors/source/solana/chain"
	"blep.ai/data/database"

	"github.com/avast/retry-go"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	Namespace = "solana"
)

type SolanaConnector struct {
	Client   *rpc.Client
	WsClient *ws.Client
	Topics   map[string]kafkautils.Topic
	KP       *kafkautils.Producer
	Conf     *viper.Viper
	Db       *database.Db
}

func NewConnector(
	client *rpc.Client,
	wsClient *ws.Client,
	topics map[string]kafkautils.Topic,
	kp *kafkautils.Producer,
	conf *viper.Viper,
	db *database.Db,
) *SolanaConnector {
	return &SolanaConnector{
		Client:   client,
		WsClient: wsClient,
		Topics:   topics,
		KP:       kp,
		Conf:     conf,
		Db:       db,
	}
}

func (c *SolanaConnector) Start(ctx context.Context) {
	sub, err := c.WsClient.RootSubscribe()
	if err != nil {
		log.Fatal().Err(err).Msg("Solana: WS Root Subscription error")
	}
	defer sub.Unsubscribe()

	sink := make(chan *kafkautils.Message, 10000)
	err = c.KP.EnableTransactions()
	if err != nil {
		log.Fatal().Err(err).Msg("Transaction was not enabled")
	}
	go c.KP.WriteAndCommitSink(sink)

	for {
		rootResult, err := sub.Recv()
		slot := uint64(*rootResult)
		if err != nil {
			log.Error().Err(err).Uint64("rootResult", slot).Msg("Failed to get root")
			continue
		}

		var out *rpc.GetBlockResult
		err = c.withRetry(
			func() error {
				var err error
				includeRewards := true
				out, err = c.Client.GetBlockWithOpts(
					ctx,
					slot,
					&rpc.GetBlockOpts{
						Commitment: "finalized",
						Rewards:    &includeRewards,
					},
				)
				return err
			},
		)

		if err != nil {
			log.Error().Err(err).Uint64("rootSlot", slot).Msg("failed to get block")
			// Store failed rpc call for later retry and backfill
			c.insertRpcError("block", fmt.Sprint(slot), "rootSlot")
			continue
		}

		c.publishBlock(out, sink)
		c.publishTransactions(out, sink)
		c.publishAccounts(ctx, out, sink)
		c.publishRewards(out, sink)
	}
}

func (c *SolanaConnector) publishBlock(out *rpc.GetBlockResult, sink chan<- *kafkautils.Message) {
	block := &chain.Block{}
	block.UnmarshalSolanaBlock(out)

	msg := &kafkautils.Message{
		Topic:    c.Topics["solanablock"],
		Key:      kafkautils.NewKey(Namespace, block.Hash),
		ProtoMsg: block,
	}

	sink <- msg
}

func (c *SolanaConnector) publishAccounts(ctx context.Context, out *rpc.GetBlockResult, sink chan<- *kafkautils.Message) {
	accounts := make([]*chain.Account, 0)

	for _, reward := range out.Rewards {
		var accountInfo *rpc.GetAccountInfoResult
		err := c.withRetry(
			func() error {
				var err error
				accountInfo, err = c.Client.GetAccountInfoWithOpts(
					ctx,
					reward.Pubkey,
					&rpc.GetAccountInfoOpts{Encoding: solana.EncodingJSONParsed},
				)
				return err
			},
		)

		if err != nil {
			log.Error().Err(err).Str("pubkey", reward.Pubkey.String()).Msg("failed to get account")
			// Store failed rpc call for later retry and backfill
			c.insertRpcError("account", reward.Pubkey.String(), "pubkey")
			continue
		}
		accountMsg := &chain.Account{}
		accountMsg.UnmarshalSolanaAccount(accountInfo.Value, reward.Pubkey)
		accounts = append(accounts, accountMsg)

		account := accountInfo.Value
		for account != nil && !account.Owner.IsZero() {
			var ownerAccountInfo *rpc.GetAccountInfoResult
			err := c.withRetry(
				func() error {
					var err error
					ownerAccountInfo, err = c.Client.GetAccountInfoWithOpts(
						ctx,
						account.Owner,
						&rpc.GetAccountInfoOpts{Encoding: solana.EncodingJSONParsed},
					)
					return err
				},
			)

			if err != nil {
				log.Error().Err(err).Str("owner pubkey", account.Owner.String()).Msg("failed to get account")
				// Store failed rpc call for later retry and backfill
				c.insertRpcError("account", account.Owner.String(), "pubkey")
				continue
			}

			accountMsg := &chain.Account{}
			accountMsg.UnmarshalSolanaAccount(ownerAccountInfo.Value, account.Owner)
			accounts = append(accounts, accountMsg)
			account = ownerAccountInfo.Value
		}
	}

	// Reverse slice so that accounts without an owner/parent are created before other accounts.
	for i, j := 0, len(accounts)-1; i < j; i, j = i+1, j-1 {
		accounts[i], accounts[j] = accounts[j], accounts[i]
	}

	for _, account := range accounts {

		msg := &kafkautils.Message{
			Topic:    c.Topics["solanaaccount"],
			Key:      kafkautils.NewKey(Namespace, out.Blockhash.String()),
			ProtoMsg: account,
		}

		sink <- msg
	}
}

func (c *SolanaConnector) publishRewards(out *rpc.GetBlockResult, sink chan<- *kafkautils.Message) {
	blockHash := out.Blockhash.String()

	for _, reward := range out.Rewards {
		blockReward := &chain.BlockReward{}
		blockReward.UnmarshalSolanaBlockReward(reward, blockHash)

		msg := &kafkautils.Message{
			Topic:    c.Topics["solanareward"],
			Key:      kafkautils.NewKey(Namespace, out.Blockhash.String()),
			ProtoMsg: blockReward,
		}

		sink <- msg
	}
}

func (c *SolanaConnector) publishTransactions(out *rpc.GetBlockResult, sink chan<- *kafkautils.Message) {
	for _, tx := range out.Transactions {
		transaction := &chain.Transaction{}
		transaction.UnmarshalSolanaTransaction(tx, out.Blockhash.String())

		msg := &kafkautils.Message{
			Topic:    c.Topics["solanatransaction"],
			Key:      kafkautils.NewKey(Namespace, out.Blockhash.String()),
			ProtoMsg: transaction,
		}

		sink <- msg

		for _, metaReward := range tx.Meta.Rewards {
			reward := &chain.BlockReward{}
			reward.UnmarshalSolanaBlockRewardForTx(metaReward, transaction)

			msg := &kafkautils.Message{
				Topic:    c.Topics["solanareward"],
				Key:      kafkautils.NewKey(Namespace, out.Blockhash.String()),
				ProtoMsg: reward,
			}

			sink <- msg
		}

		{
			message := &chain.Message{}
			message.UnmarshalSolanaMsg(tx.Transaction.Message, transaction)

			msg = &kafkautils.Message{
				Topic:    c.Topics["solanamessage"],
				Key:      kafkautils.NewKey(Namespace, out.Blockhash.String()),
				ProtoMsg: message,
			}

			sink <- msg
		}
	}
}

func (c *SolanaConnector) withRetry(retryFunc retry.RetryableFunc) error {
	maxRetries := c.Conf.GetUint("solana.maxRetries")
	retryDelay := c.Conf.GetInt("solana.retryDelay")
	opts := []retry.Option{
		retry.Delay(time.Duration(retryDelay) * time.Millisecond),
		retry.Attempts(maxRetries),
	}

	err := retry.Do(retryFunc, opts...)
	return err
}

func (c *SolanaConnector) insertRpcError(entity, id, idType string) {
	rpcErrorsTable := fmt.Sprintf(`"%s"`, c.Conf.GetString("solana.rpcErrorsTable"))
	c.Db.InsertRpcErrors(rpcErrorsTable, entity, id, idType)
}
