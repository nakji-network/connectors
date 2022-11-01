use std::time::Duration;

use crate::near_proto::access_key::{FullAccessPermission, FunctionCallPermission, Permission};
use crate::near_proto::{AccessKey, Transaction};
use crossbeam_channel::Sender;
use near_lake_framework::near_indexer_primitives::{
    views::{AccessKeyPermissionView, AccessKeyView},
    IndexerTransactionWithOutcome, StreamerMessage,
};
use prost_types::Timestamp;

use super::action_from_view;
use super::handle::ParseStreamMessage;

impl ParseStreamMessage<Transaction> for Transaction {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<Transaction>) {
        // Parse ts from timestamp_nanosec in block header since tx do not have a ts
        let ts_duration = Duration::from_nanos(message.block.header.timestamp_nanosec);
        let ts = Timestamp {
            seconds: ts_duration.as_secs() as i64,
            nanos: ts_duration.subsec_nanos() as i32,
        };
        for shard in message.shards {
            if shard.chunk.is_some() {
                let transactions = shard.chunk.unwrap().transactions;
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

impl From<IndexerTransactionWithOutcome> for Transaction {
    fn from(
        IndexerTransactionWithOutcome {
            transaction,
            outcome: _,
        }: IndexerTransactionWithOutcome,
    ) -> Self {
        Self {
            signer_id: transaction.signer_id.to_string(),
            public_key: transaction.public_key.to_string(),
            nonce: transaction.nonce,
            receiver_id: transaction.receiver_id.to_string(),
            hash: transaction.hash.to_string(),
            actions: transaction
                .actions
                .iter()
                .map(|a| action_from_view(a.to_owned()))
                .collect(),
            ts: None,
        }
    }
}

impl From<AccessKeyView> for AccessKey {
    fn from(AccessKeyView { nonce, permission }: AccessKeyView) -> Self {
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
            nonce,
            permission: Some(permission_proto),
        }
    }
}
