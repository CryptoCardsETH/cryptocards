pragma solidity ^0.4.17;

contract Battles {
	struct Battle {
		// Timestamp from the block when the battle was created
		uint64 creationTime;

		// Opponent 1 BattleGroup
		uint256 op1BattleGroup;

		// Opponent 2 BattleGroup
		uint256 op2BattleGroup;

		// Result of Battle
		uint32 battleResult;
	}
	
	// Array of all existing Battles
	Battle[] battles;

	// New Battle Event: Emitted every time a new BattleGroup is created
	event NewBattle(uint256 battleID, uint256 op1BattleGroup, uint256 op2BattleGroup);

	function createBattle(uint256 _op1, uint256 _op2) external returns (uint)
	{
		Battle memory _battle = Battle({
			creationTime: uint64(now),
			op1BattleGroup: _op1,
			op2BattleGroup: _op2,
			battleResult: 0
		});
		uint256 newID = battles.push(_battle) - 1;

		// Make sure we never overflow the battle max (4 billion battles)
		require(newID == uint256(uint32(newID)));

		// Emit NewBattleGroup Event
		NewBattle(
			newID,
			_battle.op1BattleGroup,
			_battle.op2BattleGroup
		);

		return newID;
	}

	// Return total count of all battles
	function countBattles() public view returns (uint) {
		return battles.length;
	}
}