# Nakji NEAR Connector

The NEAR protocol is a sharded, proof-of-stake, layer-one blockchain that is simple to use, secure and scalable. The Nakji NEAR Connector makes use of NEAR Lake Framework to parse events in Rust, then passes them to a Nakji connector written in Go using websockets.

## Events

As of now, the Nakji NEAR connector indexes:

* Blocks
* Transactions
* Receipts
* ExecutionOutcomes
