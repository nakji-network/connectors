# OKC smart contract Connector

Listen events from OKC smart contracts

# How to Build a Connector

## Prerequisites

1. Golang 1.18+
2. Docker
3. protoc
4. protoc-gen-go
5. abigen
6. solc

## Overview

Listen events from OKC smart contracts

### Code of a Connector

#### 1. Getting the ABI files for a smart contract

For now we just support only one ABI smart contract

#### 3. Creating the contract handling files

First, enter the <contract_name>/ directory. Create the <contract_name>.abi file in <contract_name>/ and paste the ABI
source code you copied earlier.

Next, generate the `<contract_name>.abi.go` file using `abigen` (you should have `abigen` installed from the above step)
.

**Need to add abigen and protogen installation guide links here**

```shell
abigen --abi <contract_name>.abi --pkg <contract_name> --out <contract_name>.abi.go
```

Next, generate the <contract_name>.pb.go file using `protoc`(you should have `protoc` installed from the above step)
.

```shell
protoc --go_out=. --go_opt=paths=source_relative ./<contract_name>.proto
```

```shell
go mod init

go mod tidy
```

This automatically creates the `go.mod` and `go.sum` files that tell go to fetch which modules and dependencies your
program depends on.

Next, spin up kafka by running:

```shell
docker compose up -d
```

You can update rpc in file `local.yaml` and update `manifest.yaml` for kafka transaction id

Now start the connector locally by running the `main.go` file. (If you are running to “connection refused” error, make
sure docker is running). OKC connector support backfill event log with some flag

`
--from-block(uint64): Start backfill from specific block (For the first time run backfill. It should be the block when deploy smart contract)
`

`
--num-blocks(uint64): Number of blocks you want to backfill. If you want to crawl to the newest block you can skip this flag
`


Listening new event log
```
go run cmd/okc/main.go
```

Backfill mode
You can use either or both flags
```
go run cmd/okc/main.go --from-block <blockNumber> --num-blocks <number>
```
