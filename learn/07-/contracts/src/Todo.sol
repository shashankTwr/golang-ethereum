// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Todo {

    error Todo__NotOwner();

    struct Task{
        string task;
        bool completed;
    }

    Task[] public tasks;
    address private owner;


    modifier isOwner(){
        if(msg.sender != owner) 
            revert Todo__NotOwner();
        _;
    }

    constructor(){
        owner= msg.sender;
    }

    // create Task  
    function addTask(string memory _task) public isOwner{
        Task memory task = Task({task: _task, completed: false});
        tasks.push(task);
    }

    // read task
    function readTask(uint256 _index) public view isOwner returns(Task memory todo) {
        return tasks[_index];
    }
    function readTaskContentOnIndex(uint256 _index) public view  isOwner returns(string memory task) {
        return tasks[_index].task;
    }
    function readTaskCompletedOnIndex(uint256 _index) public view isOwner  returns(bool ){
        return tasks[_index].completed;
    }


    // read complete tasks array
    function listTasks() public view  isOwner   returns(Task[] memory){
        return tasks;
    }

    
    // update Task
    function completeTask(uint256 _index) public  isOwner   {
        tasks[_index].completed = true;
    }
    function toggleTask(uint256 _index) public   isOwner  {
        tasks[_index].completed = !tasks[_index].completed;
    }


    // delete task
    function deleteTask(uint256 _index) public  isOwner   {
        for(uint256 i = _index; i < tasks.length; i++){
             tasks[i - 1] = tasks[i];
        }
        tasks.pop();
    }
    
    function getOwner() public view returns(address){
        return owner;
    }
}