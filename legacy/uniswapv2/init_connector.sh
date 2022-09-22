#!/bin/bash

# Use this command to compile abi interfaces for uniswapv2.
# Output: build_abis/*.abi.go

cd $(dirname $0)

cd smart-contracts/contracts/interfaces
solcjs --abi *.sol -o ../../../tmp_abis
cd ../../..

for f in tmp_abis/*
do
	echo $f
	TYPE=$(basename $f .abi | cut -d'_' -f3-)
	abigen --abi=$f --pkg=$(basename "$PWD") --type=$TYPE --out=$TYPE.abi.go
done

#rm -rf tmp_abis
