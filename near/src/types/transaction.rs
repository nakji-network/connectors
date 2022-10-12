use std::time::Duration;

use crate::near_proto::access_key::{FullAccessPermission, FunctionCallPermission, Permission};
use crate::near_proto::{
    action::{self, AddKey, CreateAccount, DeleteAccount, DeleteKey, FunctionCall, Stake},
    AccessKey, Action, Transaction,
};
use action::{DeployContract, Transfer};
use crossbeam_channel::Sender;
use near_lake_framework::near_indexer_primitives::{
    views::{AccessKeyPermissionView, AccessKeyView, ActionView},
    IndexerTransactionWithOutcome, StreamerMessage,
};
use prost_types::Timestamp;

use super::handle::ParseStreamMessage;

impl ParseStreamMessage<Transaction> for Transaction {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<Transaction>) {
        // Parse ts from timestamp_nanosec in block header since tx do not have a ts
        let ts_duration = Duration::from_nanos(message.block.header.timestamp_nanosec);
        let ts = Timestamp {
            seconds: ts_duration.as_secs() as i64,
            nanos: ts_duration.subsec_nanos() as i32,
        };
        for shard in &message.shards {
            if shard.chunk.is_some() {
                let transactions = &shard.chunk.as_ref().unwrap().transactions;
                for tx in transactions {
                    let mut tx_proto = Transaction::from(tx);
                    // Set transaction ts to block ts
                    tx_proto.ts = Some(ts.clone());

                    sender
                        .send(tx_proto)
                        .expect("Unable to send tx to NearHandle Sender");
                }
            }
        }
    }
}

impl From<&IndexerTransactionWithOutcome> for Transaction {
    fn from(
        IndexerTransactionWithOutcome {
            transaction,
            outcome: _,
        }: &IndexerTransactionWithOutcome,
    ) -> Self {
        let mut actions_proto = Vec::<Action>::new();

        for action in &transaction.actions {
            match action {
                ActionView::CreateAccount => actions_proto.push(Action {
                    action: Some(action::Action::CreateAccount(CreateAccount {})),
                }),
                ActionView::DeployContract { code } => actions_proto.push(Action {
                    action: Some(action::Action::DeployContract(DeployContract {
                        code: code.to_string(),
                    })),
                }),
                ActionView::FunctionCall {
                    method_name,
                    args,
                    gas,
                    deposit,
                } => actions_proto.push(Action {
                    action: Some(action::Action::FunctionCall(FunctionCall {
                        method_name: method_name.to_string(),
                        args: args.to_string(),
                        gas: *gas,
                        deposit: deposit.to_ne_bytes().to_vec(),
                    })),
                }),
                ActionView::Transfer { deposit } => actions_proto.push(Action {
                    action: Some(action::Action::Transfer(Transfer {
                        deposit: deposit.to_ne_bytes().to_vec(),
                    })),
                }),
                ActionView::Stake { stake, public_key } => actions_proto.push(Action {
                    action: Some(action::Action::Stake(Stake {
                        stake: stake.to_ne_bytes().to_vec(),
                        public_key: public_key.to_string(),
                    })),
                }),
                ActionView::AddKey {
                    public_key,
                    access_key,
                } => actions_proto.push(Action {
                    action: Some(action::Action::AddKey(AddKey {
                        public_key: public_key.to_string(),
                        access_key: Some(AccessKey::from(access_key)),
                    })),
                }),
                ActionView::DeleteKey { public_key } => actions_proto.push(Action {
                    action: Some(action::Action::DeleteKey(DeleteKey {
                        public_key: public_key.to_string(),
                    })),
                }),
                ActionView::DeleteAccount { beneficiary_id } => actions_proto.push(Action {
                    action: Some(action::Action::DeleteAccount(DeleteAccount {
                        beneficiary_id: beneficiary_id.to_string(),
                    })),
                }),
            }
        }

        Self {
            signer_id: transaction.signer_id.to_string(),
            public_key: transaction.public_key.to_string(),
            nonce: transaction.nonce,
            receiver_id: transaction.receiver_id.to_string(),
            hash: transaction.hash.to_string(),
            actions: actions_proto,
            ts: None,
        }
    }
}

impl From<&AccessKeyView> for AccessKey {
    fn from(AccessKeyView { nonce, permission }: &AccessKeyView) -> Self {
        let permission_proto: Permission;
        match permission {
            AccessKeyPermissionView::FunctionCall {
                allowance,
                receiver_id,
                method_names,
            } => {
                permission_proto = Permission::FunctionCall(FunctionCallPermission {
                    allowance: allowance.map(|x| x.to_ne_bytes().to_vec()),
                    receiver_id: receiver_id.to_string(),
                    method_names: method_names.to_vec(),
                })
            }
            AccessKeyPermissionView::FullAccess => {
                permission_proto = Permission::FullAccess(FullAccessPermission {})
            }
        }
        Self {
            nonce: *nonce,
            permission: Some(permission_proto),
        }
    }
}
