use std::thread;

use crossbeam_channel::Receiver;

use prost::Message;
use ws::listen;

use crate::near_proto::NearMessage;

pub struct NearWsServer {
    pub listen_addr: String,
    pub channel: Receiver<NearMessage>,
}

impl NearWsServer {
    pub fn start(&self) {
        println!("Listening for connections on {}", self.listen_addr);
        // Begin NEAR Stream over ws on connection
        listen(self.listen_addr.clone(), |out| {
            // Sender for messages to be broadcast over ws
            let out_stream = out.clone();
            // Receiver for messages from handle_streamer_message
            let channel = self.channel.clone();
            thread::spawn(move || loop {
                match channel.recv() {
                    Ok(msg) => {
                        let msg_serial = Self::serialize(msg);
                        out_stream.broadcast(msg_serial).expect("Unable to broadcast message to websockets.");
                    }
                    Err(_) => {
                        println!("Unable to receive messages. Terminating.");
                        break;
                    }
                }
            });

            // Handle messages from client
            move |msg| {
                Ok({
                    println!("Server got message {:?}. ", msg);
                })
            }
        })
        .expect("Failed to start ws server");
    }

    fn serialize(message: NearMessage) -> Vec<u8> {
        let mut buf = Vec::new();
        buf.reserve(message.encoded_len());

        message.encode(&mut buf).expect("Unable to serialize proto message");

        buf
    }
}
