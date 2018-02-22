pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Battles.sol";

contract TestBattles {
	Battles battles = Battles(DeployedAddresses.Battles());

	// Testing the createBattle() function
	function testCreateBattle() public {
		uint expected = battles.countBattles();

		uint256 op1 = uint256(1);
		uint256 op2 = uint256(2);

		uint returnedID = battles.createBattle(op1, op2);

		Assert.equal(returnedID, expected, "New battle should have next sequential battleID");
	}

	// Test creating many additional battles
	function testCreateManyBattles() public {
		uint numBattles = 5;
		uint expected = battles.countBattles() + numBattles;

		uint i = 0;
		for (i = 0; i < numBattles; i++) {
			testCreateBattle();
		}

		uint256 totalBattles = battles.countBattles();
		Assert.equal(totalBattles, expected, "Count of battles should increase by numBattles");
	}

}