pub use crossbeam_channel::Sender;
use futures::StreamExt;
use near_lake_framework::{
    near_indexer_primitives::StreamerMessage, LakeConfig, LakeConfigBuilder,
};
use std::{
    fs::{self, read_to_string, File},
    io::Write,
};

#[tokio::main]
pub async fn start_lake_stream(
    stream_sender: Sender<StreamerMessage>,
    from_block: u64,
    num_blocks: u64,
) {
    // create a NEAR Lake Framework config
    let config = lake_config(from_block);

    println!("",);

    // instantiate the NEAR Lake Framework Stream
    let (_sender, stream) = near_lake_framework::streamer(config);

    // read the stream events and pass them to a handler function with
    // buffered list of pending futures (size 4)
    let mut handlers = tokio_stream::wrappers::ReceiverStream::new(stream)
        .map(|streamer_message| handle_streamer_message(streamer_message, stream_sender.clone()))
        .buffer_unordered(4usize);

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

fn lake_config(from_block: u64) -> LakeConfig {
    // Create a NEAR Lake Framework config
    // Begin by checking config volume for credentials
    let yaml_string = read_to_string("/etc/nakji/config.yaml").ok();

    // If credentials are in config directory, we move them to where aws expects them to be
    if yaml_string.is_some() {
        let yaml = yaml_rust::YamlLoader::load_from_str(&yaml_string.unwrap())
            .expect("Found config file at /etc/nakji/config.yaml, but unable to parse it");
        // Multi-document support. We are only reading 1
        let doc = &yaml[0];
        // Read AWS ID and Secret
        let aws_id = doc["aws"]["id"].as_str().unwrap();
        let aws_secret = doc["aws"]["secret"].as_str().unwrap();

        let home_dir = home::home_dir()
            .expect("Unable to find home dir")
            .to_str()
            .expect("Unable to get home dir string")
            .to_string();

        // Add aws directory to directory path
        let mut aws_dir = home_dir.clone();
        aws_dir.push_str("/.aws");

        fs::create_dir(aws_dir.clone()).expect("Could not create directory for credentials");

        // Create str representation of credentials file
        let mut credential_str = "[default]\naws_access_key_id=".to_string();
        credential_str.push_str(aws_id);
        credential_str.push_str("\naws_secret_access_key=");
        credential_str.push_str(aws_secret);

        // Add credentials file to directory path
        let mut credentials_dir = aws_dir.clone();
        credentials_dir.push_str(&"/credentials".to_string());

        // Write to file
        let mut file = File::create(credentials_dir).expect("Unable to create credentials file");
        file.write_all(credential_str.as_bytes())
            .expect("Unable to write credentials str to file");

        LakeConfigBuilder::default()
            .mainnet()
            .start_block_height(from_block)
            .blocks_preload_pool_size(2000)
            .build()
            .expect("Unable to find AWS credentials. To automatically configure, add them to ~/.aws/credentials")
    } else {
        println!("Unable to find AWS credentials on config volume. Checking for credentials in local environment.");
        LakeConfigBuilder::default()
            .mainnet()
            .start_block_height(from_block)
            .blocks_preload_pool_size(2000)
            .build()
            .expect("Unable to find AWS credentials. To automatically configure, add them to ~/.aws/credentials")
    }
}
