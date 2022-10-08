use std::thread;

use clap::{Parser, Subcommand};
use crossbeam_channel::unbounded;
use nakji_near_client::{
    lake_stream,
    near_proto::{Block, Transaction},
    types::handle::ParseStreamMessage,
    ws_server,
};
use near_jsonrpc_client::{methods, JsonRpcClient};
use near_primitives::types::{BlockReference, Finality};
use ws_server::NearWsServer;

#[derive(Parser, Debug, Clone)]
#[clap(version = "0.0.1", author = "Nakji")]
pub(crate) struct Opts {
    #[clap(subcommand)]
    pub start_opts: StartOptions,
    port: u16,
}

#[derive(Subcommand, Debug, Clone)]
pub(crate) enum StartOptions {
    FromBlock { height: u64 },
    FromLatest,
}

impl Opts {
    pub fn start_options(&self) -> &StartOptions {
        &self.start_opts
    }
}

fn main() {
    let opts = Opts::parse();

    let rt = tokio::runtime::Runtime::new().unwrap();
    let future = get_start_block_height(&opts);
    let start_block = rt.block_on(future);

    // lake_stream_send: Sender that NEAR Lake stream will write to
    // lake_stream_rx: Receiver for StreamerMessages from NEAR Lake
    let (lake_stream_send, lake_stream_rx) = unbounded();

    // Create a vector of join handles
    // If the join handles are not held, the threads will detach, causing a resource leak
    let mut join_handles = Vec::new();

    // Start lake stream with output to lake_stream_send
    println!("Starting NEAR Lake Stream");
    join_handles.push(thread::spawn(move || {
        lake_stream::start_lake_stream(lake_stream_send, start_block)
    }));

    // Create sender and receiver for NEAR Handlers
    let (event_send, event_rx) = unbounded();

    // ws config
    let mut ws_host = "127.0.0.1:".to_string();
    ws_host.push_str(&opts.port.to_string());

    // Create ws servers for each event type. Send near handler output (receiver) to ws.
    let server = NearWsServer {
        listen_addr: ws_host.clone(),
        channel: event_rx,
    };
    join_handles.push(thread::spawn(move || server.start()));

    /*
    For each message received from NEAR Lake Framework stream, borrow and handle message,
    then send resulting proto message
    */
    loop {
        match lake_stream_rx.recv() {
            Ok(stream_msg) => {
                Block::handle_streamer_message(&stream_msg, event_send.clone());
                Transaction::handle_streamer_message(&stream_msg, event_send.clone());
            }
            Err(_) => {
                println!(
                    "Unable to receive messages from Near Lake Framework Stream. Terminating."
                );
                break;
            }
        }
    }

    for handle in join_handles {
        handle
            .join()
            .expect("Couldn't join on the associated thread");
    }
}

async fn get_start_block_height(opts: &Opts) -> u64 {
    match opts.start_options() {
        StartOptions::FromBlock { height } => *height,
        StartOptions::FromLatest => final_block_height("https://rpc.mainnet.near.org").await,
    }
}

async fn final_block_height(rpc_url: &str) -> u64 {
    println!("Getting latest block height");
    let client = JsonRpcClient::connect(rpc_url);
    let request = methods::block::RpcBlockRequest {
        block_reference: BlockReference::Finality(Finality::None),
    };

    let latest_block = client.call(request).await.unwrap();

    latest_block.header.height
}
