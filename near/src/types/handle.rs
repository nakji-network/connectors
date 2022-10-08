use crossbeam_channel::{Receiver, Sender};
use near_lake_framework::near_indexer_primitives::StreamerMessage;

use crate::near_proto::NearMessage;

/*
    This was used when there was a ws endpoint for each event type. Each event type had a handler for managing
    senders and receivers. This is not necessary currently, as all events are wrapped in a single protobuf
    message type. It may become necessary to have handler types again in the future, so I'm leaving this 
    here for now. 
*/
pub struct NearHandler {
    pub receiver: Receiver<StreamerMessage>,
    pub sender: Sender<NearMessage>,
}

impl NearHandler {
    pub fn new_handler(receiver: Receiver<StreamerMessage>, sender: Sender<NearMessage>) -> NearHandler {
        Self { receiver, sender }
    }
}

pub trait ParseStreamMessage<T> {
    fn handle_streamer_message(message: &StreamerMessage, sender: Sender<NearMessage>);
}
