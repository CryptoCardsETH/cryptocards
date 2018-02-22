pragma solidity ^0.4.17;

contract Battle {
	
}

contract BattleGroups {
	struct BattleGroup {
		// The timestamp from the block when the battlegroup was created
		uint64 creationTime;

		// Address of owner of BattleGroup
		address owner;

		// Cards which belong to the BattleGroup
		uint256[3] cards;
	}

	// Array of all existing GameCards
	BattleGroup[] battlegroups;

	// New BattleGroup Event: Emitted every time a new BattleGroup is created
	event NewBattleGroup(address owner, uint256 battleGroupID, uint256[3] cards);

	function createBattleGroup(address _owner, uint256[3] _cards) external returns (uint)
	{
		BattleGroup memory _group = BattleGroup({
			creationTime: uint64(now),
			owner: _owner,
			cards: _cards
		});
		uint256 newID = battlegroups.push(_group) - 1;

		// Make sure we never overflow the battlegroup max (4 billion cards)
		require(newID == uint256(uint32(newID)));

		// Emit NewBattleGroup Event
		NewBattleGroup(
			_owner,
			newID,
			_group.cards
		);

		return newID;
	}

	// Return total count of all battleGroups
	function countBattleGroups() public view returns (uint) {
		return battlegroups.length;
	}
}