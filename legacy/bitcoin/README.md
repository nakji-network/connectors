# Running BTC Connector

We will need `btcd` installed with the following steps in the command line:

    go get github.com/btcsuite/btcd
    cd $GOPATH/src/github.com/btcsuite/btcd
    GO111MODULE=on go install -v . ./cmd/...

Make sure that go binaries in your PATH with `export PATH="$HOME/go/bin:$PATH"` and maybe add that to
bashrc/zshrc.

Once `btcd` is installed, you may run the following commands for the BTC connector to work:

    make start-btcd
    make start-bitcoin


Note that if you see a line like `Catching up indexes from height 181186 to 373507` when starting `btcd`,
that means the RPC serve for `btcd` has not started yet and will only be started after all blocks have
been indexed. Once all are indexed, the RPC server will be started and `make start-bitcoin` will be
able to connect its RPC client to `btcd`.

## Running BTC Connector in Prod

We will need to set `bitcoin.rpc.url` key in analytics-config secret for prod to `"btcd:8334"`.