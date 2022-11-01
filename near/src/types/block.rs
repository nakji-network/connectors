use std::time::Duration;

use crate::near_proto::{
    block::{block_header::SlashedValidator, BlockHeader, Chunk, ValidatorStake},
    Block,
};
use crossbeam_channel::Sender;
use near_lake_framework::near_indexer_primitives::{
    views::{
        validator_stake_view::ValidatorStakeView, BlockHeaderView, BlockView, ChunkHeaderView,
    },
    StreamerMessage,
};
use prost_types::Timestamp;

use super::handle::ParseStreamMessage;

impl ParseStreamMessage<Block> for Block {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<Block>) {
        sender
            .send(Block::from(&message.block))
            .expect("Unable to send block to NearHandle Sender");
    }
}

impl From<&BlockView> for Block {
    fn from(
        BlockView {
            author,
            header,
            chunks,
        }: &BlockView,
    ) -> Self {
        let mut chunks_proto = Vec::<Chunk>::new();
        for chunk in chunks {
            let chunk_proto = Chunk::from(chunk);
            chunks_proto.push(chunk_proto);
        }

        // Parse ts from timestamp_nanosec
        let ts_duration = Duration::from_nanos(header.timestamp_nanosec);
        let ts = Timestamp {
            seconds: ts_duration.as_secs() as i64,
            nanos: ts_duration.subsec_nanos() as i32,
        };

        Self {
            author: author.to_string(),
            header: Some(BlockHeader::from(header)),
            chunks: chunks_proto,
            ts: Some(ts),
        }
    }
}

impl From<&BlockHeaderView> for BlockHeader {
    fn from(
        BlockHeaderView {
            height,
            prev_height,
            epoch_id,
            next_epoch_id,
            hash,
            prev_hash,
            prev_state_root,
            chunk_receipts_root,
            chunk_headers_root,
            chunk_tx_root,
            outcome_root,
            chunks_included,
            challenges_root,
            timestamp,
            timestamp_nanosec,
            random_value,
            validator_proposals,
            chunk_mask,
            gas_price,
            block_ordinal,
            rent_paid: _,
            validator_reward: _,
            total_supply,
            challenges_result,
            last_final_block,
            last_ds_final_block,
            next_bp_hash,
            block_merkle_root,
            epoch_sync_data_hash,
            approvals,
            signature,
            latest_protocol_version,
        }: &BlockHeaderView,
    ) -> Self {
        // Get vector of String representation of approval hashes
        let mut approvals_str = Vec::<String>::new();
        for approval in approvals {
            if approval.is_some() {
                approvals_str.push(approval.as_ref().unwrap().to_string())
            }
        }

        let mut slashed_validators = Vec::<SlashedValidator>::new();
        for challenge in challenges_result {
            let mut slashed_validator = SlashedValidator::default();
            slashed_validator.account_id = challenge.account_id.to_string();
            slashed_validator.is_double_sign = challenge.is_double_sign;
            slashed_validators.push(slashed_validator);
        }

        let validator_proposals = parse_validator_proposals(validator_proposals.to_vec());

        Self {
            epoch_id: epoch_id.to_string(),
            next_epoch_id: next_epoch_id.to_string(),
            hash: hash.to_string(),
            prev_hash: prev_hash.to_string(),
            prev_state_root: prev_state_root.to_string(),
            chunk_receipts_root: chunk_receipts_root.to_string(),
            chunk_headers_root: chunk_headers_root.to_string(),
            chunk_tx_root: chunk_tx_root.to_string(),
            outcome_root: outcome_root.to_string(),
            chunks_included: *chunks_included,
            challenges_root: challenges_root.to_string(),
            random_value: random_value.to_string(),
            timestamp: *timestamp,
            timestamp_nanosec: *timestamp_nanosec,
            chunk_mask: chunk_mask.to_vec(),
            gas_price: gas_price.to_ne_bytes().to_vec(),
            block_ordinal: *block_ordinal,
            total_supply: total_supply.to_ne_bytes().to_vec(),
            last_final_block: last_final_block.to_string(),
            last_ds_final_block: last_ds_final_block.to_string(),
            next_bp_hash: next_bp_hash.to_string(),
            block_merkle_root: block_merkle_root.to_string(),
            epoch_sync_data_hash: epoch_sync_data_hash.map(|h| h.to_string()),
            signature: signature.to_string(),
            latest_protocol_version: *latest_protocol_version,
            height: *height,
            prev_height: *prev_height,
            approvals: approvals_str,
            challenges_result: slashed_validators,
            validator_proposals,
        }
    }
}

impl From<&ChunkHeaderView> for Chunk {
    fn from(
        ChunkHeaderView {
            chunk_hash,
            prev_block_hash,
            outcome_root,
            prev_state_root,
            encoded_merkle_root,
            encoded_length,
            height_created,
            height_included,
            shard_id,
            gas_used,
            gas_limit,
            rent_paid: _,
            validator_reward: _,
            balance_burnt,
            outgoing_receipts_root,
            tx_root,
            validator_proposals,
            signature,
        }: &ChunkHeaderView,
    ) -> Self {
        let validator_proposals = parse_validator_proposals(validator_proposals.to_vec());

        Self {
            chunk_hash: chunk_hash.to_string(),
            prev_block_hash: prev_block_hash.to_string(),
            outcome_root: outcome_root.to_string(),
            prev_state_root: prev_state_root.to_string(),
            encoded_merkle_root: encoded_merkle_root.to_string(),
            encoded_length: *encoded_length,
            height_created: *height_created,
            height_included: *height_included,
            shard_id: *shard_id,
            gas_used: *gas_used,
            gas_limit: *gas_limit,
            balance_burnt: balance_burnt.to_ne_bytes().to_vec(),
            outgoing_receipts_root: outgoing_receipts_root.to_string(),
            tx_root: tx_root.to_string(),
            validator_proposals,
            signature: signature.to_string(),
        }
    }
}

fn parse_validator_proposals(validator_proposals: Vec<ValidatorStakeView>) -> Vec<ValidatorStake> {
    let mut validator_stakes = Vec::<ValidatorStake>::new();

    for stake_view in validator_proposals {
        match stake_view {
            ValidatorStakeView::V1(stake_view) => validator_stakes.push(ValidatorStake {
                account_id: stake_view.account_id.to_string(),
                public_key: stake_view.public_key.to_string(),
                stake: stake_view.stake.to_ne_bytes().to_vec(),
            }),
        }
    }

    validator_stakes
}
