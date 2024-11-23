// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {Todo} from "../src/Todo.sol";

contract DeployTodo is Script {
    Todo public todo;

    function run() external returns(Todo) {
        vm.startBroadcast();
        todo = new Todo();
        vm.stopBroadcast();
        return todo;
    }
}
