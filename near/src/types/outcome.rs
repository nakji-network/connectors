use std::time::Duration;

use crossbeam_channel::Sender;
use near_lake_framework::near_indexer_primitives::views::{
    ExecutionOutcomeView, ExecutionOutcomeWithIdView, ExecutionStatusView,
};
use near_lake_framework::near_indexer_primitives::StreamerMessage;

use near_primitives::errors::{ActionError, InvalidTxError};
use near_primitives::errors::{ActionErrorKind, TxExecutionError};
use prost_types::Timestamp;

use crate::near_proto::execution_outcome::outcome::{
    failure, Failure, Status, SuccessReceiptId, SuccessValue, Unknown,
};
use crate::near_proto::execution_outcome::MerklePath;
use crate::near_proto::{execution_outcome::Outcome, ExecutionOutcome};

use super::handle::ParseStreamMessage;

impl ParseStreamMessage<ExecutionOutcome> for ExecutionOutcome {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<ExecutionOutcome>) {
        // Parse ts from timestamp_nanosec in block header since execution outcomes do not have a ts
        let ts_duration = Duration::from_nanos(message.block.header.timestamp_nanosec);
        let ts = Some(Timestamp {
            seconds: ts_duration.as_secs() as i64,
            nanos: ts_duration.subsec_nanos() as i32,
        });
        for shard in message.shards {
            if shard.chunk.is_some() {
                let transactions = shard.chunk.unwrap().transactions;
                for tx in transactions {
                    let mut outcome_proto = ExecutionOutcome::from(tx.outcome.execution_outcome);
                    outcome_proto.ts = ts.clone();

                    sender
                        .send(outcome_proto)
                        .expect("Unable to send execution outcome to NearHandle Sender");
                }
            }
            for outcome in shard.receipt_execution_outcomes {
                let mut outcome_proto = ExecutionOutcome::from(outcome.execution_outcome);
                outcome_proto.ts = ts.clone();

                sender
                    .send(outcome_proto)
                    .expect("Unable to send execution outcome to NearHandle Sender");
            }
        }
    }
}

impl From<ExecutionOutcomeWithIdView> for ExecutionOutcome {
    fn from(
        ExecutionOutcomeWithIdView {
            proof,
            block_hash,
            id,
            outcome,
        }: ExecutionOutcomeWithIdView,
    ) -> Self {
        Self {
            ts: None,
            id: id.to_string(),
            block_hash: block_hash.to_string(),
            outcome: Some(Outcome::from(outcome)),
            proof: proof
                .iter()
                .map(|p| MerklePath {
                    hash: p.hash.to_string(),
                    direction: match p.direction {
                        near_primitives::merkle::Direction::Left => 1,
                        near_primitives::merkle::Direction::Right => 2,
                    },
                })
                .collect(),
        }
    }
}

impl From<ExecutionOutcomeView> for Outcome {
    fn from(
        ExecutionOutcomeView {
            logs,
            receipt_ids,
            gas_burnt,
            tokens_burnt,
            executor_id,
            status,
            metadata: _,
        }: ExecutionOutcomeView,
    ) -> Self {
        let status_proto = match status {
            ExecutionStatusView::Unknown => Some(Status::Unknown(Unknown {})),
            ExecutionStatusView::Failure(f) => match f {
                TxExecutionError::ActionError(e) => Some(Status::from(e)),
                TxExecutionError::InvalidTxError(e) => Some(Status::from(e)),
            },
            ExecutionStatusView::SuccessValue(v) => {
                Some(Status::SuccessValue(SuccessValue { value: v }))
            }
            ExecutionStatusView::SuccessReceiptId(i) => {
                Some(Status::SuccessReceiptId(SuccessReceiptId {
                    receipt_id: i.to_string(),
                }))
            }
        };

        Self {
            executor_id: executor_id.to_string(),
            logs,
            receipt_ids: receipt_ids.iter().map(|r| r.to_string()).collect(),
            gas_burnt: gas_burnt.to_ne_bytes().to_vec(),
            tokens_burnt: tokens_burnt.to_ne_bytes().to_vec(),
            status: status_proto,
        }
    }
}

impl From<ActionError> for Status {
    fn from(ActionError { index, kind }: ActionError) -> Self {
        match kind {
            ActionErrorKind::AccountAlreadyExists { account_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 1,
                })),
            }),
            ActionErrorKind::AccountDoesNotExist { account_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 2,
                })),
            }),
            ActionErrorKind::CreateAccountOnlyByRegistrar {
                account_id: _,
                registrar_account_id: _,
                predecessor_id: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 3,
                })),
            }),
            ActionErrorKind::CreateAccountNotAllowed {
                account_id: _,
                predecessor_id: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 4,
                })),
            }),
            ActionErrorKind::ActorNoPermission {
                account_id: _,
                actor_id: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 5,
                })),
            }),
            ActionErrorKind::DeleteKeyDoesNotExist {
                account_id: _,
                public_key: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 6,
                })),
            }),
            ActionErrorKind::AddKeyAlreadyExists {
                account_id: _,
                public_key: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 7,
                })),
            }),
            ActionErrorKind::DeleteAccountStaking { account_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 8,
                })),
            }),
            ActionErrorKind::LackBalanceForState {
                account_id: _,
                amount: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 9,
                })),
            }),
            ActionErrorKind::TriesToUnstake { account_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 10,
                })),
            }),
            ActionErrorKind::TriesToStake {
                account_id: _,
                stake: _,
                locked: _,
                balance: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 11,
                })),
            }),
            ActionErrorKind::InsufficientStake {
                account_id: _,
                stake: _,
                minimum_stake: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 12,
                })),
            }),
            ActionErrorKind::FunctionCallError(_) => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 13,
                })),
            }),
            ActionErrorKind::NewReceiptValidationError(_) => Self::Failure(Failure {
                error: Some(failure::Error::ActionError(failure::ActionError {
                    index,
                    kind: 14,
                })),
            }),
            ActionErrorKind::OnlyImplicitAccountCreationAllowed { account_id: _ } => {
                Self::Failure(Failure {
                    error: Some(failure::Error::ActionError(failure::ActionError {
                        index,
                        kind: 15,
                    })),
                })
            }
            ActionErrorKind::DeleteAccountWithLargeState { account_id: _ } => {
                Self::Failure(Failure {
                    error: Some(failure::Error::ActionError(failure::ActionError {
                        index,
                        kind: 16,
                    })),
                })
            }
        }
    }
}

impl From<InvalidTxError> for Status {
    fn from(e: InvalidTxError) -> Self {
        match e {
            InvalidTxError::InvalidAccessKeyError(_) => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(1)),
            }),
            InvalidTxError::InvalidSignerId { signer_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(2)),
            }),
            InvalidTxError::SignerDoesNotExist { signer_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(3)),
            }),
            InvalidTxError::InvalidNonce {
                tx_nonce: _,
                ak_nonce: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(4)),
            }),
            InvalidTxError::NonceTooLarge {
                tx_nonce: _,
                upper_bound: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(5)),
            }),
            InvalidTxError::InvalidReceiverId { receiver_id: _ } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(6)),
            }),
            InvalidTxError::InvalidSignature => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(7)),
            }),
            InvalidTxError::NotEnoughBalance {
                signer_id: _,
                balance: _,
                cost: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(8)),
            }),
            InvalidTxError::LackBalanceForState {
                signer_id: _,
                amount: _,
            } => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(9)),
            }),
            InvalidTxError::CostOverflow => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(10)),
            }),
            InvalidTxError::InvalidChain => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(11)),
            }),
            InvalidTxError::Expired => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(12)),
            }),
            InvalidTxError::ActionsValidation(_) => Self::Failure(Failure {
                error: Some(failure::Error::InvalidTxError(13)),
            }),
            InvalidTxError::TransactionSizeExceeded { size: _, limit: _ } => {
                Self::Failure(Failure {
                    error: Some(failure::Error::InvalidTxError(14)),
                })
            }
        }
    }
}
