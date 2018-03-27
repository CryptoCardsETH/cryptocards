pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol"; 
import "../contracts/CardBase.sol"; 
import "../contracts/Blacklist.sol"; 

contract TestBlacklist {

  function testAddToBlacklist() public {
    Blacklist blacklist = new Blacklist();  
    bool notAddedYet = false;
    bool isAdded = blacklist.isOnBlacklist(tx.origin); 
    Assert.equal(notAddedYet,isAdded, "User should not be on the blacklist yet"); 

    blacklist.AddToBlacklist(tx.origin); 
    isAdded = blacklist.isOnBlacklist(tx.origin); 
    notAddedYet = true; 
    Assert.equal(notAddedYet, isAdded, "User should now be on the blacklist"); 
  }

  function testRemoveFromBlacklist() public {

    Blacklist blacklist = new Blacklist(); 

    blacklist.AddToBlacklist(tx.origin); 
    bool isAdded = blacklist.isOnBlacklist(tx.origin); 
    bool notAddedYet = true; 
    Assert.equal(notAddedYet, isAdded, "User should now be on the blacklist"); 

    blacklist.removeFromBlacklist(tx.origin); 
    notAddedYet = false;
    isAdded = blacklist.isOnBlacklist(tx.origin); 
    Assert.equal(notAddedYet,isAdded, "User should not be on the blacklist now"); 

  }
}
