pragma solidity ^0.4.17; 

contract Blacklist {

  mapping(address => bool) public blacklist;

  function AddToBlacklist(address _user) public {
    blacklist[_user] = true; 
  }

  function isOnBlacklist(address _user) public view returns (bool) {
    bool status = blacklist[_user]; 
    return status;  
  }

  function removeFromBlacklist(address _user) public returns (bool) {
    if (blacklist[_user]) {
      blacklist[_user] = false; 
      return true;  
    }
    return false; 
  }
}
