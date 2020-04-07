pragma solidity ^0.5.12;

import { ERC20Mintable } from "../../../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";
import { IERC20 } from "../../../node_modules/openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import { CustomIncrementCoinageMock as CustomIncrementCoinage } from "../../../node_modules/coinage-token/flatten.sol";

import { RootChainI } from "../../RootChainI.sol";

import { RootChainRegistryI } from "./RootChainRegistryI.sol";
import { DepositManagerI } from "./DepositManagerI.sol";
import { PowerTONI } from "./PowerTONI.sol";


interface GovI {
  function changeGov(address uint256) external view;

}