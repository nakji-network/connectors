pub mod types;
pub mod ws_server;
pub mod lake_stream;

// Include the `near` module, which is generated from near.proto
pub mod near_proto {
    include!(concat!(env!("OUT_DIR"), "/near.rs"));
}
