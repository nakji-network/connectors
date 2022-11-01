use std::thread;

use crossbeam_channel::Receiver;

use prost::Message;
use ws::listen;

use crate::near_proto::{Block, ExecutionOutcome, Receipt, Transaction};

pub struct NearWsServer {
    pub listen_addr: String,
    pub block_rx: Receiver<Block>,
    pub tx_rx: Receiver<Transaction>,
    pub receipt_rx: Receiver<Receipt>,
    pub outcome_rx: Receiver<ExecutionOutcome>,
}

pub struct WsRouter {
    pub listen_addr: String,
    pub block_rx: Receiver<Block>,
    pub tx_rx: Receiver<Transaction>,
    pub receipt_rx: Receiver<Receipt>,
    pub outcome_rx: Receiver<ExecutionOutcome>,
    sender: ws::Sender,
    inner: Box<dyn ws::Handler>,
}

impl ws::Handler for WsRouter {
    fn on_request(&mut self, req: &ws::Request) -> ws::Result<ws::Response> {
        // Clone the sender so that we can move it into the child handler
        let out = self.sender.clone();

        match req.resource() {
            // Route to a data handler
            "/block" => {
                self.inner = Box::new(NearWsEndpoint::<Block> {
                    ws: out,
                    rx: self.block_rx.clone(),
                })
            }

            // Route to a data handler
            "/tx" => {
                self.inner = Box::new(NearWsEndpoint::<Transaction> {
                    ws: out,
                    rx: self.tx_rx.clone(),
                })
            }

            // Route to a data handler
            "/receipt" => {
                self.inner = Box::new(NearWsEndpoint::<Receipt> {
                    ws: out,
                    rx: self.receipt_rx.clone(),
                })
            }

            // Route to a data handler
            "/outcome" => {
                self.inner = Box::new(NearWsEndpoint::<ExecutionOutcome> {
                    ws: out,
                    rx: self.outcome_rx.clone(),
                })
            }

            // Use the default child handler, NotFound
            _ => (),
        }

        // Delegate to the child handler
        self.inner.on_request(req)
    }

    // Pass through any other methods that should be delegated to the child.

    fn on_shutdown(&mut self) {
        self.inner.on_shutdown()
    }

    fn on_open(&mut self, shake: ws::Handshake) -> ws::Result<()> {
        self.inner.on_open(shake)
    }

    fn on_message(&mut self, msg: ws::Message) -> ws::Result<()> {
        self.inner.on_message(msg)
    }

    fn on_close(&mut self, code: ws::CloseCode, reason: &str) {
        self.inner.on_close(code, reason)
    }

    fn on_error(&mut self, err: ws::Error) {
        self.inner.on_error(err);
    }
}

struct NearWsEndpoint<T: Message> {
    ws: ws::Sender,
    rx: Receiver<T>,
}

impl<T: Message + 'static> ws::Handler for NearWsEndpoint<T> {
    fn on_open(self: &mut NearWsEndpoint<T>, _shake: ws::Handshake) -> ws::Result<()> {
        // Receiver for messages from handle_streamer_message
        let channel = self.rx.clone();
        let ws = self.ws.clone();
        thread::spawn(move || loop {
            match channel.recv() {
                Ok(msg) => {
                    let mut buf = Vec::new();
                    buf.reserve(msg.encoded_len());

                    msg.encode(&mut buf)
                        .expect("Unable to serialize proto message");

                    ws.send(buf)
                        .expect("Unable to broadcast message to websockets.");
                }
                Err(_) => {
                    println!("Source channel closed. Terminating.");
                    ws.close(ws::CloseCode::Normal).expect("Unable to close client ws connection");
                    break;
                }
            }
        });
        Ok(())
    }
}

// This handler returns a 404 response to all handshake requests
struct NotFound;

impl ws::Handler for NotFound {
    fn on_request(&mut self, req: &ws::Request) -> ws::Result<ws::Response> {
        // This handler responds to all requests with a 404
        let mut res = ws::Response::from_request(req)?;
        res.set_status(404);
        res.set_reason("Not Found");
        Ok(res)
    }
}

impl NearWsServer {
    pub fn start(&self) {
        // Begin NEAR Stream over ws on connection
        listen(self.listen_addr.clone(), |out| WsRouter {
            sender: out,
            inner: Box::new(NotFound),
            listen_addr: self.listen_addr.clone(),
            block_rx: self.block_rx.clone(),
            tx_rx: self.tx_rx.clone(),
            receipt_rx: self.receipt_rx.clone(),
            outcome_rx: self.outcome_rx.clone(),
        })
        .expect("Unable to start ws server.");
    }
}
