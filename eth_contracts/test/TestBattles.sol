pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CryptoCardsCore.sol";

contract TestBattles {
	CryptoCardsCore cryptoCards = CryptoCardsCore(DeployedAddresses.CryptoCardsCore());

	function testBattlesContract() public {
		Battles _battleContract = cryptoCards.BattleContract();
		Assert.notEqual(address(_battleContract), 0, "Battle Contract Should Exist");

		BattleGroups _battleGroupContract = cryptoCards.BattleGroupContract();
		Assert.notEqual(address(_battleGroupContract), 0, "Battle Group Contract Should Exist");
	}

	// Testing the createBattleGroup() function
	function testCreateBattleGroup() public returns (uint256) {
		uint expected = cryptoCards.BattleGroupContract().countBattleGroups();

		uint256 card1 = cryptoCards.createGen0Card();
		uint256 card2 = cryptoCards.createGen0Card();
		uint256 card3 = cryptoCards.createGen0Card();
		uint256 card4 = cryptoCards.createGen0Card();
		uint256 card5 = cryptoCards.createGen0Card();
		uint256[5] memory cards = [card1, card2, card3, card4, card5];

		uint returnedID = cryptoCards.BattleGroupContract().createBattleGroup(this, cards);
		Assert.equal(returnedID, expected, "New BattleGroup should have next sequential battleGroupID");

		return returnedID;
	}

	// Testing the createBattle() function
	function testCreateBattle() public returns (uint256) {
		uint expected = cryptoCards.BattleContract().countBattles();

		uint256 group1 = testCreateBattleGroup();
		uint256 group2 = testCreateBattleGroup();

		uint returnedID = cryptoCards.BattleContract().createBattle(group1, group2);
		Assert.equal(returnedID, expected, "New battle should have next sequential battleID");

		// Send the same groups to battle again - should fail as on cooldown
		uint returnedID2 = cryptoCards.BattleContract().createBattle(group1, group2);
		Assert.equal(returnedID2, 0, "Second battle while on cooldown should fail");

		return returnedID;
	}

	// Test creating many additional battles
	function testCreateManyBattles() public {
		uint numBattles = 4;
		uint expected = cryptoCards.BattleContract().countBattles() + numBattles;

		uint i = 0;
		for (i = 0; i < numBattles; i++) {
			testCreateBattle();
		}

		uint256 totalBattles = cryptoCards.BattleContract().countBattles();
		Assert.equal(totalBattles, expected, "Count of battles should increase by numBattles");
	}
}