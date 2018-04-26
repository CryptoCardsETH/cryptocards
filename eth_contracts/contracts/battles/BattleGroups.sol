pragma solidity ^0.4.17;
import "../CryptoCardsCore.sol";

contract BattleGroups {
	uint public constant MAX_CARDS_PER_GROUP = 5;

	struct BattleGroup {
		uint64 creationTime;	// The timestamp from the block when the battlegroup was created
		address owner;			// Address of owner of BattleGroup
		uint256[5] cards;		// Cards which belong to the BattleGroup
	}

	// Array of all existing BattleGroups
	BattleGroup[] battlegroups;

	// New BattleGroup Event: Emitted every time a new BattleGroup is created
	event NewBattleGroup(address owner, uint256 battleGroupID, uint256[5] cards);

	CryptoCardsCore public cryptoCardsContract; // Pointer to Card Contract
	function BattleGroups() public {
		CryptoCardsCore candidateContract = CryptoCardsCore(msg.sender);
		//require(candidateContract.supportsInterface(InterfaceSignature_ERC721));
		cryptoCardsContract = candidateContract;
	}

	function createBattleGroup(address _owner, uint256[5] _cards) external returns (uint256) {
		BattleGroup memory _group = BattleGroup({
			creationTime: uint64(now),
			owner: _owner,
			cards: _cards
		});
		uint256 newID = battlegroups.push(_group) - 1;

		// Make sure we never overflow the battlegroup max (4 billion groups)
		require(newID == uint256(uint32(newID)));

		// Emit NewBattleGroup Event
		NewBattleGroup(
			_owner,
			newID,
			_group.cards
		);

		return newID;
	}

	function isGroupReadyForBattle(uint256 groupID) public view returns (bool) {
		require (groupID < countBattleGroups());
		BattleGroup memory _group = battlegroups[groupID];

		uint i = 0;
		for (i = 0; i < _group.cards.length; i++) {
			if (!cryptoCardsContract.isReadyForBattle(_group.cards[i])) {
				return false;
			}
		}
		return true;
	}

	function setGroupOnBattleCooldown(uint256 groupID) public returns (bool) {
		require (groupID < countBattleGroups());
		BattleGroup memory _group = battlegroups[groupID];

		uint i = 0;
		for (i = 0; i < _group.cards.length; i++) {
			cryptoCardsContract.setOnBattleCooldown(_group.cards[i]);
		}
		return true;
	}

	// Return total count of all battleGroups
	function countBattleGroups() public view returns (uint) {
		return battlegroups.length;
	}
}