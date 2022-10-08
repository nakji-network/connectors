pub use crossbeam_channel::Sender;
use futures::StreamExt;
use near_lake_framework::{near_indexer_primitives::StreamerMessage, LakeConfigBuilder};

#[tokio::main]
pub async fn start_lake_stream(stream_sender: Sender<StreamerMessage>, start_block: u64) {
    // create a NEAR Lake Framework config
    let config = LakeConfigBuilder::default()
        .mainnet()
        .start_block_height(start_block)
        .build()
        .expect("Failed to build LakeConfig");

    // instantiate the NEAR Lake Framework Stream
    let (_sender, stream) = near_lake_framework::streamer(config);

    // read the stream events and pass them to a handler function with
    // concurrency 1
    let mut handlers = tokio_stream::wrappers::ReceiverStream::new(stream)
        .map(|streamer_message| handle_streamer_message(streamer_message, stream_sender.clone()))
        .buffer_unordered(1usize);

    while let Some(_handle_message) = handlers.next().await {}
    drop(handlers); // close the channel so the sender will stop
}

// The handler function to take the entire `StreamerMessage`
// and print the block height and number of shards
async fn handle_streamer_message(
    streamer_message: near_lake_framework::near_indexer_primitives::StreamerMessage,
    sender: Sender<StreamerMessage>,
) {
    println!("Received message from NEAR Lake");
    sender.send(streamer_message).unwrap();
}
