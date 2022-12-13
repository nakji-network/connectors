# Klaytn Connector

This is the example connector used at https://docs.nakji.network.

## Klaytn
Klaytn is an open-source public blockchain for all who wish to build, work, or play in the metaverse.

## How to run

### Prerequisites
- Kafka
- Update local.yaml with appropriate values
- Protobuf (Install Protobuf for Golang by following this: https://developers.google.com/protocol-buffers/docs/gotutorial)

### local.yaml example
```shell
log:
 level: debug
 format: console
kafka:
 url: localhost:9092
 env: dev
rpcs:
 opb:
  full:
    - your infura websocket endpoint here
```
### Run protobuf complier
```shell
# To runs protobuf compiler
make proto
```
### Run the connector
```shell
# To runs connector
go run cmd/klaytn/main.go
```