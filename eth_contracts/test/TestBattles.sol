pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CryptoCardsCore.sol";

contract TestBattles {
	CryptoCardsCore cryptoCards = CryptoCardsCore(DeployedAddresses.CryptoCardsCore());

	//////////////////// Contract Tests ////////////////////

	function testBattlesContract() public {
		Battles _battleContract = cryptoCards.BattleContract();
		Assert.notEqual(address(_battleContract), 0, "Battle Contract Should Exist");

		BattleGroups _battleGroupContract = cryptoCards.BattleGroupContract();
		Assert.notEqual(address(_battleGroupContract), 0, "Battle Group Contract Should Exist");

		BattleQueue _battleQueueContract = cryptoCards.BattleQueueContract();
		Assert.notEqual(address(_battleQueueContract), 0, "Battle Queue Contract Should Exist");
	}

	//////////////////// Battle Groups ////////////////////

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

	//////////////////// Battles ////////////////////

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

	//////////////////// Battle Queue ////////////////////

	function testBattleQueue() public returns (bool) {
		uint expected = cryptoCards.BattleContract().countBattles() + 1;

		uint256 group1 = testCreateBattleGroup();
		uint256 group2 = testCreateBattleGroup();

		bool joinSuccess = cryptoCards.BattleQueueContract().joinQueue(group1);
		Assert.equal(joinSuccess, true, "First Battle Group Joining Queue Failed");

		bool joinSuccess2 = cryptoCards.BattleQueueContract().joinQueue(group2);
		Assert.equal(joinSuccess2, true, "Second Battle Group Joining Queue Failed");

		uint256 totalBattles = cryptoCards.BattleContract().countBattles();
		Assert.equal(totalBattles, expected, "Count of battles should increase by numBattles");
	}

	// Test creating many additional battles
	function testCreateManyQueues() public {
		uint numQueues = 3;
		uint expected = cryptoCards.BattleContract().countBattles() + numQueues;

		uint i = 0;
		for (i = 0; i < numQueues; i++) {
			testBattleQueue();
		}

		uint256 totalBattles = cryptoCards.BattleContract().countBattles();
		Assert.equal(totalBattles, expected, "Count of battles should increase by numQueues");
	}
}