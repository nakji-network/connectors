# Nakji NEAR Connector

The NEAR protocol is a sharded, proof-of-stake, layer-one blockchain that is simple to use, secure and scalable. The Nakji NEAR Connector makes use of NEAR Lake Framework to parse events in Rust, then passes them to a Nakji connector written in Go using websockets (no RPC involved, other than finding most recent block height).

## Events

As of now, the Nakji NEAR connector indexes:

* Blocks
* Transactions
  
Note: work is underway to add support for `Receipts` and `ExecutionOutcomes`, with `Receipts`, `Transactions`, and `ExecutionOutcomes` all being separate events with separate tables etc.

## Rust Connector

The Rust portion of the Nakji connector streams data from the NEAR chain using NEAR Lake Framework, transforms this data to Prost types using Nakji's Protobuf representation of NEAR events, serializes the data, and broadcasts it over websockets. I had originally used different endpoints for different events, but fixing issues with routing was taking too long. So now all events are wrapped in one message type.

### Development Environment

It is recommended that devlopers use rustup to initialize their Rust environment. Installing Rust directly with Brew can cause issues with IDE extensions, like rust-analyzer for VS Code. Develop in VS Code for cool features like Rust types auto updating with near.proto, useful errors/suggestions/type info, auto installation of dependencies + more. 

#### Setup

* Install rustup: `brew install rustup`
* Initialize: `rustup-init`
* Install rust-analyzer to make your life so much easier

#### Run

* Run for development: `cargo run main.rs`
* Build for production: `cargo build --release` (binary will be at `target/release/nakji_near_client`)
* If you want to manually run for local development along with the Go connector, consider commenting out the portion of the connect function in client.go that executes the Rust binary.

### Adding Events

To add events, add to the proto file, create a new file in `src/types`, implement `From` for the necessary NEAR primitives, and implement the `ParseStreamMessage` trait for the generated type to handle `StreamerMessage` that comes from NEAR Lake Framework. Finally, follow the pattern of existing types in `main.rs` to create a `NearHandler` type, create a `NearWsServer`, and handle incoming messages in the main loop using the `handle_streamer_message` function you implemented. Note: In order for the new events to be sent to Kafka, `client.go` will also need to be updated (see below).

### Future Improvements

* Many NEAR types do not implement `Copy` or `Clone`, meaning that events are parsed using a borrowed reference to a `StreamerMessage`. This means that only one event type is parsed at a time. 
* Backfill is inconsistent with other connectors. Currently, backfill is not separated from other events. 
* Similarly, the connector supports `from-block` but not `num-blocks`.
* All events are sent to every client. Adding filtering on the rust side would be nice.
* Consider integrating into chain?

## Go Connector

The Go portion of the connector is very similar to other connectors, other than `client.go`, which defines ws endpoint and event type mappings, generates ws listeners for each endpoint, and sends events to an event chan. Since all events are received as proto messages and require no filtering, the connector writes events from this chan straight to `EventSink`. 

### Adding Events

Assuming you've already taken care of things on the Rust side, all you need to do is add to the eventTypes map in the `NewConnector` function in `near.go`. A ws listener will automatically be generated. If you are attempting to build a separate connector that uses a subset of NEAR events, just create a new connector with a different map in `NewConnector`. Should you want to filter/manipulate events, consider creating a function that uses events received from `Client.events` before passing them to `EventSink`.

