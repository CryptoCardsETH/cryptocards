pragma solidity ^0.4.17;
//import "../CryptoCardsCore.sol";

contract Battles {
	struct Battle {
		uint64 creationTime; // Timestamp from the block when the battle was created
		uint256 op1BattleGroup; // Opponent 1 BattleGroup
		uint256 op2BattleGroup; // Opponent 2 BattleGroup
		uint32 battleResult; // Result of Battle
	}
	
	Battle[] battles; // Array of all existing Battles

	// New Battle Event: Emitted every time a new BattleGroup is created
	event BattleResult(uint256 battleID, uint256 winnerBattleGroup, uint256 loserBattleGroup);

	//CryptoCardsCore public cryptoCardsContract; // Pointer to Card Contract
	function Battles() public {
		//CryptoCardsCore candidateContract = CryptoCardsCore(msg.sender);
        //require(candidateContract.supportsInterface(InterfaceSignature_ERC721));
        //cryptoCardsContract = candidateContract;
	}

	function createBattle(uint256 _op1, uint256 _op2) public returns (uint) {
		Battle memory _battle = Battle({
			creationTime: uint64(now),
			op1BattleGroup: _op1,
			op2BattleGroup: _op2,
			battleResult: 0
		});
		uint256 newID = battles.push(_battle) - 1;

		// Make sure we never overflow the battle max (4 billion battles)
		require(newID == uint256(uint32(newID)));

		// Emit Battle Result Event
		BattleResult(
			newID,
			_battle.op1BattleGroup,
			_battle.op2BattleGroup
		);

		return newID;
	}

	function countBattles() public view returns (uint) { // Return total count of all battles
		return battles.length;
	}
}