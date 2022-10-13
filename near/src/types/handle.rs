use std::thread;

use crossbeam_channel::{Receiver, Sender, unbounded};
use near_lake_framework::near_indexer_primitives::{StreamerMessage, views::BlockView, IndexerShard};
use prost::Message;

/*
    This was used when there was a ws endpoint for each event type. Each event type had a handler for managing
    senders and receivers. This is not necessary currently, as all events are wrapped in a single protobuf
    message type. It may become necessary to have handler types again in the future, so I'm leaving this 
    here for now. 
*/
pub struct NearHandler<T: Message + ParseStreamMessage<T> + 'static> {
    stream_rx: Receiver<Vec<u8>>,
    sender: Sender<T>,
    pub rx: Receiver<T>
}

impl<T: Message + ParseStreamMessage<T>> NearHandler<T> {
    pub fn new_handler(stream_rx: Receiver<Vec<u8>>) -> NearHandler<T> {
        let (sender, rx) = unbounded();
        Self { stream_rx, sender, rx }
    }

    pub fn start(&self) -> thread::JoinHandle<()> {
        let lake_stream_rx = self.stream_rx.clone();
        let sender = self.sender.clone();
        //let sender = self.sender.clone();
        thread::spawn(move || loop {
            match lake_stream_rx.recv() {
                Ok(stream_msg) => {
                    //Block::handle_streamer_message(&stream_msg, event_send.clone());
                    //Transaction::handle_streamer_message(&stream_msg, event_send.clone());
                    let stream_message_decode: StreamerMessage = serde_json::from_slice(&stream_msg).expect("Handler unable to decode message");
                    <T>::handle_streamer_message(stream_message_decode, sender.clone())
                }
                Err(_) => {
                    println!(
                        "Unable to receive messages from Near Lake Framework Stream. Terminating."
                    );
                    break;
                }
            }
        })
    }
}

pub trait ParseStreamMessage<T> {
    fn handle_streamer_message(message: StreamerMessage, sender: Sender<T>);
}
