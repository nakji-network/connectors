# WooFi

This connector ingests the woofi contracts specified below. Related connectors are in ../woofi_avax, ../woofi_bsc, ../woofi_fantom, ../woofi_polygon.

ABIs are based on the ones [provided by the Woofi Team](https://github.com/woonetwork/woofi_subgraph/tree/main/abis), with their names changed slightly to make it backwards compatible with the currently deployed connector + its users.

## Address <> ABI mappings | start

-- BSC --

0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F | 13662625 | ABI: WooPPV3

0xcEf5BE73ae943B77f9Bc08859367D923C030a269 | 17015221 | ABI: WooRouterV2

0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83 | 11983352 | ABI: WooRouterV1

0x53E255e8Bbf4EDF16797f9885291B3Ca0C70B59f | 18675185 | ABI: WooCrossChainRouterV1

0x8489d142Da126F4Ea01750e80ccAa12FD1642988 | 12094867 | ABI: WooPPV2

0x10C24658815585851a8888f059Cb4338790023F1 | 11981813 | ABI: WooPPV2

-- Avalanche --

0x1df3009c57a8B143c6246149F00B090Bce3b8f88 | 13965201 | ABI: WooPPV3

0x5AA6a4E96A9129562e2fc06660D07FEdDAAf7854 | 13966541 | ABI: WooRouterV2

0x3e0Da0A9e4139B32b37710784b8DCa643c152001 | 13522432 | ABI: WooRouterV2

0x160020B09DeD3d862f7f851B5c50632BcF2062FF | 10857815 | ABI: WooRouterV2

0xdF37F7A85D4563f39A78494568824b4dF8669B7a | 16011735 | ABI: WooCrossChainRouterV1

0xF8cE0D043891b62c55380fB1EFBfB4F186153D96 | 9713408 | ABI: WooPPV3

-- Polygon --

0x7400B665C8f4f3a951a99f1ee9872efb8778723d | 29860023 | ABI: WooPPV3

0x9D1A92e601db0901e69bd810029F2C14bCCA3128 | 29860251 | ABI: WooRouterV2

0x376d567C5794cfc64C74852A9DB2105E0b5B482C | 29897168 | ABI: WooCrossChainRouterV1

-- Fantom --

0x9503E7517D3C5bc4f9E4A1c6AE4f8B33AC2546f2 | 35475008 | ABI: WooPPV3

0x37B5a5A730dAD670874f26Cc5507bb1b9705e447 | 35476301 | ABI: WooRouterV2

0xcF6Ce5Fd6bf28bB1AeAc88A55251f6c840059De5 | 40483282 | ABI: WooCrossChainRouterV1

## Naming scheme

The ABI names are based on backwards compatibility and contract names.

subgraph/WooPPV1_1.json               => WooPPV2.abi

subgraph/WooPPV1_2.json               => WooPPV3.abi

subgraph/WooRouterV1_1.json           => WooRouterV1.abi

subgraph/WooRouterV1_2.json           => WooRouterV2.abi

subgraph/WooCrossChainRouterV1_1.json => WooCrossChainRouterV1.abi
