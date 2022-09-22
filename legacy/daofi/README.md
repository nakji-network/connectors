# Daofi Connector
## Daofi
A decentralized exchange that allows for fully customizable bonding curves

1. Ensure uniswapv2 git submodule is in smart-contracts.
2. `./compile_abis.sh` to generate files from smart contracts in submodule.

Generated files are *.abi.go



### old notes
subscribe to new tokens added to uniswap. Or from an open source github repo.

grab token addresss from array

subscribe to eth logs for all token addresses

parse log using abi

case: each function call type


{"jsonrpc":"2.0", "id": 1, "method": "eth_subscribe", "params": ["logs", {"address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"}]}


usdc-eth pair 0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc


{"jsonrpc":"2.0", "id": 1, "method": "eth_subscribe", "params": ["logs", {"address": "0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"}]}
https://goethereumbook.org/event-subscribe/

abigen --abi=connectors/source/uniswapv2/uniswapv2pair.abi --pkg=uniswapv2 --out=connectors/source/uniswapv2/uniswapv2pair.go
to generate go file from abi

