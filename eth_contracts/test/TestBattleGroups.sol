pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Battle.sol";

contract TestBattleGroups {
  BattleGroups battleGroups = BattleGroups(DeployedAddresses.BattleGroups());

	// Testing the createBattleGroup() function
	function testCreateBattleGroup() public {
		uint expected = battleGroups.countBattleGroups();

		uint256[3] memory cards = [uint256(1), uint256(2), uint256(3)];

		uint returnedID = battleGroups.createBattleGroup(this, cards);

		Assert.equal(returnedID, expected, "Created card should have next sequential cardID");
	}

	// Test creating many additional battlegroups
	function testCreateManyCards() public {
		uint numGroups = 5;
		uint expected = battleGroups.countBattleGroups() + numGroups;

		uint256[3] memory cards = [uint256(1), uint256(2), uint256(3)];

		uint i = 0;
		for (i = 0; i < numGroups; i++) {
			battleGroups.createBattleGroup(this, cards);
		}

		uint256 totalGroups = battleGroups.countBattleGroups();
		Assert.equal(totalGroups, expected, "Count of battlegroups should increase by numGroups");
	}

}