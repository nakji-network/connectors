pragma solidity ^0.5.17;

import "./erc20/ERC20Detailed.sol";
import "./erc20/ERC20Mintable.sol";


contract Roninweth is ERC20Detailed, ERC20Mintable {
  constructor () ERC20Detailed("Ronin Wrapped Ether", "WETH", 18)
    public
  {}
}