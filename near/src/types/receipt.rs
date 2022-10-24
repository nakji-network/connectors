use std::time::Duration;

use crossbeam_channel::Sender;
use near_lake_framework::near_indexer_primitives::views::{
    DataReceiverView, ReceiptEnumView, ReceiptView,
};
use near_lake_framework::near_indexer_primitives::StreamerMessage;
use prost_types::Timestamp;

use crate::near_proto::receipt;
use crate::near_proto::receipt::{ActionReceipt, DataReceipt};
use crate::near_proto::Receipt;

use super::action_from_view;
use super::handle::ParseStreamMessage;

impl ParseStreamMessage<Receipt> for Receipt {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<Receipt>) {
        // Parse ts from timestamp_nanosec in block header since receipts do not have a ts
        let ts_duration = Duration::from_nanos(message.block.header.timestamp_nanosec);
        let ts = Some(Timestamp {
            seconds: ts_duration.as_secs() as i64,
            nanos: ts_duration.subsec_nanos() as i32,
        });
        for shard in message.shards {
            if shard.chunk.is_some() {
                let chunk = shard.chunk.unwrap();
                for tx in chunk.transactions {
                    if tx.outcome.receipt.is_some() {
                        let receipt = tx.outcome.receipt.unwrap();
                        let mut receipt_proto = Receipt::from(receipt);
                        receipt_proto.ts = ts.clone();

                        sender
                            .send(receipt_proto)
                            .expect("Unable to send receipt to NearHandle Sender");
                    }
                }
                for receipt in chunk.receipts {
                    let mut receipt_proto = Receipt::from(receipt);
                    receipt_proto.ts = ts.clone();

                    sender
                        .send(receipt_proto)
                        .expect("Unable to send receipt to NearHandle Sender");
                }
            }
            for outcome in shard.receipt_execution_outcomes {
                let mut receipt_proto = Receipt::from(outcome.receipt);
                receipt_proto.ts = ts.clone();

                sender
                    .send(receipt_proto)
                    .expect("Unable to send receipt to NearHandle Sender");
            }
        }
    }
}

impl From<ReceiptView> for Receipt {
    fn from(
        ReceiptView {
            predecessor_id,
            receiver_id,
            receipt_id,
            receipt,
        }: ReceiptView,
    ) -> Self {
        let receipt_proto = match receipt {
            ReceiptEnumView::Action {
                signer_id,
                signer_public_key,
                gas_price,
                output_data_receivers,
                input_data_ids,
                actions,
            } => receipt::Receipt::ActionReceipt(ActionReceipt {
                signer_id: signer_id.to_string(),
                signer_public_key: signer_public_key.to_string(),
                gas_price: gas_price.to_ne_bytes().to_vec(),
                output_data_receivers: output_data_receivers
                    .iter()
                    .map(|r| receipt::action_receipt::DataReceiver::from(r.clone()))
                    .collect(),
                input_data_ids: input_data_ids.iter().map(|h| h.to_string()).collect(),
                actions: actions
                    .iter()
                    .map(|a| action_from_view(a.to_owned()))
                    .collect(),
            }),
            ReceiptEnumView::Data { data_id, data } => receipt::Receipt::DataReceipt(DataReceipt {
                data_id: data_id.to_string(),
                data,
            }),
        };
        Self {
            predecessor_id: predecessor_id.to_string(),
            receiver_id: receiver_id.to_string(),
            receipt_id: receipt_id.to_string(),
            receipt: Some(receipt_proto),
            ts: None,
        }
    }
}

impl From<DataReceiverView> for receipt::action_receipt::DataReceiver {
    fn from(
        DataReceiverView {
            data_id,
            receiver_id,
        }: DataReceiverView,
    ) -> Self {
        Self {
            data_id: data_id.to_string(),
            receiver_id: receiver_id.to_string(),
        }
    }
}
