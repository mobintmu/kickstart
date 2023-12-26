// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity >=0.5.0 <0.9.0;
// linter warnings (red underline) about pragma version can igonored!

// contract code will go here

contract Campaing {

    address public manager;
    uint256 public minimumContribution;
    mapping (address => bool ) public approvers;
    uint256 public approversCount;

    struct Request {
        string description;
        uint256 value;
        address payable recipient;
        bool complete;
        uint256 approvalCount;
        mapping(address => bool) approvals;
    }

    uint256 public numRequests;
    mapping(uint256 => Request) public requests;


    constructor(uint256 min ) {
        manager = msg.sender;
        minimumContribution = min;
    }

    function contibute() public payable {
        require(
            msg.value >= minimumContribution,
            "minimum contribute is required."
        );

        approvers[msg.sender] = true;
        approversCount ++;

    }

    function isPaied() public view returns (bool) {
        return (approvers[msg.sender]);
    }

    function getBalance() public view returns(uint256){
        return (address(this).balance);
    }


    function addRequest(
        string memory description,
        uint256 value,
        address payable  recipient
    ) public onlyManager {

        Request storage req = requests[numRequests++];
        req.description = description;
        req.value = value;
        req.recipient = recipient;
        req.complete = false;
        req.approvalCount = 0;
    }


    function approveRequest(uint256 index) public{

        Request storage req = requests[index];
        require(
            approvers[msg.sender] ,
            "Only contributors can approve a specific payment request"
        );

        require(
            !req.approvals[msg.sender],
            "You have already voted to approve this request"
        );

        require(index < numRequests,
        "You can vote only in valid index. more than valid");

        require(index >= 0,
        "You can vote only in valid index. minus value");

        req.approvals[msg.sender] = true;
        req.approvalCount++;
    }

    function finalizeRequest(uint256 index) public onlyManager {
        require(index < numRequests,
        "You can vote only in valid index. more than valid");

        require(index >= 0,
        "You can vote only in valid index. minus value");

        Request storage req = requests[index];
        
        require(!req.complete ,
        "This request has been completed.");

        require(
            req.approvalCount > (approversCount / 2),
            "This request needs more approvals before it can be finalized"
        );

        req.recipient.transfer(req.value);
        req.complete = true;
    }



    modifier onlyManager() {
        require(
            msg.sender == manager,
            "Only the campaign manager can call this function."
        );
        _;
    }
 
}