// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {Todo} from "src/Todo.sol";
import {DeployTodo} from "script/DeployTodo.s.sol";

contract TodoTest is Test {
    Todo todo;

    uint256 public constant STARTING_BALANCE = 1000 ether;
    address USER = makeAddr("user");

    function setUp() public {
        DeployTodo deployTodo = new DeployTodo();
        vm.deal(USER, STARTING_BALANCE);
        todo = deployTodo.run();
    }

    function test_addTask() public {
        todo.addTask("test");
        console.log(todo.getOwner());
        assertEq(todo.readTask(0).task, "test");
    }

    function test_readTask() public {
        todo.addTask("test");
        assertEq(todo.readTask(0).task, "test");
    }

    function testDeleteTask() public {
        // add few tasks
        // delete some task
        todo.addTask("test1");
        todo.addTask("test2");
        todo.addTask("test3");
        todo.addTask("test4");
        todo.addTask("test5");
        todo.addTask("test6");
        todo.addTask("test7");
        todo.addTask("test8");
        todo.addTask("test9");

        todo.deleteTask(4);

        // test3 should be deleted
        assertEq(todo.listTasks().length,8);
        assertEq(todo.readTask(4).task, "test6");
    }

    function testCompletedTasks() public {

        // completed tasks
        todo.addTask("test1");
        todo.addTask("test2");
        todo.completeTask(1);
        assertEq(todo.readTask(1).completed, true);
    }
}
