pub use crossbeam_channel::Sender;
use futures::StreamExt;
use near_lake_framework::{
    near_indexer_primitives::StreamerMessage, LakeConfig, LakeConfigBuilder,
};
use std::fs::read_to_string;

#[tokio::main]
pub async fn start_lake_stream(
    stream_sender: Sender<StreamerMessage>,
    from_block: u64,
    num_blocks: u64,
    buffer_size: usize,
    block_pool_size: usize
) {
    // create a NEAR Lake Framework config
    let config = lake_config(from_block, block_pool_size);

    // instantiate the NEAR Lake Framework Stream
    let (_sender, stream) = near_lake_framework::streamer(config);

    // read the stream events and pass them to a handler function with
    // buffered list of pending futures
    let mut handlers = tokio_stream::wrappers::ReceiverStream::new(stream)
        .map(|streamer_message| handle_streamer_message(streamer_message, stream_sender.clone()))
        .buffer_unordered(buffer_size);

    // Ensure that stream stops if we are backfilling
    if num_blocks > 0 {
        let mut block_count = 0;
        while let Some(_handle_message) = handlers.next().await {
            if !(block_count < num_blocks) {
                break;
            };
            block_count += 1;
        }
    } else {
        while let Some(_handle_message) = handlers.next().await {}
    }

    drop(handlers); // close the channel so the sender will stop
}

// Take `StreamerMessage` and send it to the given channel to be parsed
async fn handle_streamer_message(
    streamer_message: near_lake_framework::near_indexer_primitives::StreamerMessage,
    sender: Sender<StreamerMessage>,
) {
    sender.send(streamer_message).unwrap();
}

fn lake_config(from_block: u64, block_pool_size: usize) -> LakeConfig {
    // Create a NEAR Lake Framework config
    // Begin by checking config volume for credentials
    let yaml_string = read_to_string("/etc/nakji/config.yaml").ok();

    
    let yaml = yaml_rust::YamlLoader::load_from_str(&yaml_string.unwrap())
        .expect("Found config file, but unable to parse it");
    // Multi-document support. We are only reading 1
    let doc = &yaml[0];
    // Read AWS ID and Secret
    let aws_id = doc["aws"]["id"].as_str().unwrap();
    let aws_secret = doc["aws"]["secret"].as_str().unwrap();

    let credentials = aws_types::Credentials::new(
        aws_id,
        aws_secret,
        None,
        None,
        "custom_credentials",
    );
    let s3_config = aws_sdk_s3::Config::builder()
        .credentials_provider(credentials)
        .region(aws_types::region::Region::new("eu-central-1"))
        .build();
    
    // Build and return lake config
    LakeConfigBuilder::default()
        .s3_config(s3_config)
        .mainnet()
        .start_block_height(from_block)
        .blocks_preload_pool_size(block_pool_size)
        .build()
        .expect("Unable to find AWS credentials.")
    
}
