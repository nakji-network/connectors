
# MakerDAO

MakerDAO is a lending protocol. It accepts various tokens as collateral and lends Dai stablecoin against that collateral. 

Its active contracts can be found at the link below

https://changelog.makerdao.com/releases/mainnet/active/contracts.json

### Ilk Registry (Vault Management)

MakerDAO deploys new smart contracts for each vault it manages. This is done through Ilk Registry. 

There is no common smart contract for all vaults. Each vault might have its own properties, hence emit different events. So whenever there is a new smart contract, its events should be extracted. 

As a simple example, **UNI** contracts are shown below

"UNI": "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", -- Vault contract **(Gem)**
"PIP_UNI": "0xf363c7e351C96b910b92b45d34190650df4aE8e7", --	Collateral price **(Pip)**
"MCD_JOIN_UNI_A": "0x3BC3A58b4FC1CbE7e98bB4aB7c99535e8bA9b8F1",	--	Deposit/Withdraw collateral
"MCD_CLIP_UNI_A": "0x3713F83Ee6D138Ce191294C131148176015bC29a",	--	Liquidate collateral **(Clip)**
"MCD_CLIP_CALC_UNI_A": "0xeA7FE6610e6708E2AFFA202948cA19ace3F580AE",	--	Calculate how much collateral to liquidate

Among these contracts, Gem, Pip and Clip contracts emit events.
