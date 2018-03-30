pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/BattleGroups.sol";
import "../contracts/Blacklist.sol"; 

contract TestBattleGroups {
  BattleGroups battleGroups = BattleGroups(DeployedAddresses.BattleGroups());
  Blacklist blacklist = Blacklist(DeployedAddresses.Blacklist()); 

  // Testing the createBattleGroup() function
  function testCreateBattleGroup() public {

    battleGroups.setBlacklist(blacklist);  
    uint expected = battleGroups.countBattleGroups();

    uint256[5] memory cards = [uint256(1), uint256(2), uint256(3), uint256(4), uint256(5)];

    uint returnedID = battleGroups.createBattleGroup(this, cards);

    Assert.equal(returnedID, expected, "New BattleGroup should have next sequential battleGroupID");
  }

  // Test creating many additional battlegroups
  function testCreateManyGroups() public {

    battleGroups.setBlacklist(blacklist);  
    uint numGroups = 5;
    uint expected = battleGroups.countBattleGroups() + numGroups;

    uint i = 0;
    for (i = 0; i < numGroups; i++) {
      testCreateBattleGroup();
    }

    uint256 totalGroups = battleGroups.countBattleGroups();
    Assert.equal(totalGroups, expected, "Count of battlegroups should increase by numGroups");
  }
}
