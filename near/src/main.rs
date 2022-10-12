use std::{thread, process};

use clap::{Parser, Subcommand};
use crossbeam_channel::unbounded;
use nakji_near_client::{
    lake_stream,
    near_proto::{Block, Transaction},
    types::handle::NearHandler,
    ws_server,
};
use near_jsonrpc_client::{methods, JsonRpcClient};
use near_primitives::types::{BlockReference, Finality};
use ws_server::NearWsServer;

// Arg parsing

#[derive(Parser, Debug, Clone)]
pub(crate) struct Opts {
    #[clap(subcommand)]
    pub start_opts: StartOptions,
    #[clap(short, long)]
    port: u16,
}

#[derive(Subcommand, Debug, Clone)]
pub(crate) enum StartOptions {
    FromBlock { height: u64 },
    NumBlocks { num_blocks: u64 },
    FromLatest,
}

impl Opts {
    pub fn start_options(&self) -> &StartOptions {
        &self.start_opts
    }
}

fn main() {
    let opts = Opts::parse();

    // Tokio runtime for async call in synchronous function
    let rt = tokio::runtime::Runtime::new().unwrap();
    let future = get_start_opts(&opts);
    let (from_block, num_blocks) = rt.block_on(future);

    // lake_stream_send: Sender that NEAR Lake stream will write to
    // lake_stream_rx: Receiver for StreamerMessages from NEAR Lake
    let (lake_stream_send, lake_stream_rx) = unbounded();

    // Create a vector of join handles
    // If the join handles are not held, the threads will detach, causing a resource leak
    let mut join_handles = Vec::new();

    // Start lake stream with output to lake_stream_send
    join_handles.push(thread::spawn(move || {
        lake_stream::start_lake_stream(lake_stream_send, from_block, num_blocks)
    }));

    // Create senders and receivers for NEAR Handlers
    let (block_send, block_rx) = unbounded();
    let block_handler = NearHandler::<Block>::new_handler(block_rx);
    join_handles.push(block_handler.start());

    let (tx_send, tx_rx) = unbounded();
    let tx_handler = NearHandler::<Transaction>::new_handler(tx_rx);
    join_handles.push(tx_handler.start());

    // ws config
    let mut ws_host = "127.0.0.1:".to_string();
    ws_host.push_str(&opts.port.to_string());

    // Create ws server with receiver for each message type.
    let server = NearWsServer {
        listen_addr: ws_host.clone(),
        block_rx: block_handler.rx.clone(),
        tx_rx: tx_handler.rx.clone(),
    };
    join_handles.push(thread::spawn(move || server.start()));

    // For each message received from NEAR Lake Framework stream, serialize, clone, and handle
    loop {
        match lake_stream_rx.recv() {
            Ok(stream_msg) => {
                // StreamerMessage does not derive Clone, so we must convert it to a type that does
                // Fortunately, it uses serde, so we can serialize, clone, then deserialize.
                let msg_json = serde_json::to_vec(&stream_msg)
                    .expect("Unable to serialize message from lake stream");

                // Send cloned messages to handlers to be parsed
                block_send
                    .send(msg_json.clone())
                    .expect("Unable to send message from Lake Stream to Block handler");
                tx_send
                    .send(msg_json.clone())
                    .expect("Unable to send message from Lake Stream to Transaction handler");
            }
            Err(_) => {
                println!("Near Lake Framework Stream closed. Terminating.");
                break;
            }
        }
    }

    process::exit(0x0100);
}

async fn get_start_opts(opts: &Opts) -> (u64, u64) {
    match opts.start_options() {
        StartOptions::FromBlock { height } => (*height, final_block_height().await - *height),
        StartOptions::FromLatest => (final_block_height().await, 0),
        StartOptions::NumBlocks { num_blocks } => {
            let latest_block = final_block_height().await;
            (latest_block - *num_blocks, *num_blocks)
        }
    }
}

async fn final_block_height() -> u64 {
    let rpc_url = "https://rpc.mainnet.near.org";
    let client = JsonRpcClient::connect(rpc_url);
    let request = methods::block::RpcBlockRequest {
        block_reference: BlockReference::Finality(Finality::None),
    };

    let latest_block = client.call(request).await.expect("Unable to connect to NEAR RPC");

    latest_block.header.height
}
