pragma solidity ^0.4.17;
import "../CryptoCardsCore.sol";

contract Battles {
	struct Battle {
		uint64 creationTime;		// Timestamp from the block when the battle was created
		uint256 op1BattleGroup;		// Opponent 1 BattleGroup
		uint256 op2BattleGroup;		// Opponent 2 BattleGroup
		uint256 winnerBattleGroup;	// Winner BattleGroup
		uint256 loserBattleGroup;	// Winner BattleGroup
	}
	
	Battle[] battles; // Array of all existing Battles

	// New Battle Event: Emitted every time a new BattleGroup is created
	event BattleResult(uint256 battleID, uint256 winnerBattleGroup, uint256 loserBattleGroup);

	CryptoCardsCore public cryptoCardsContract; // Pointer to Card Contract
	function Battles() public {
		CryptoCardsCore candidateContract = CryptoCardsCore(msg.sender);
		//require(candidateContract.supportsInterface(InterfaceSignature_ERC721));
		cryptoCardsContract = candidateContract;
	}

	function createBattle(uint256 _op1, uint256 _op2) public returns (uint256) {
		// Ensure Ready For Battle
		if (!isReadyForBattle(_op1) || !isReadyForBattle(_op2)) {
			return 0;
		}

		// Create Battle
		Battle memory _battle = Battle({
			creationTime: uint64(now),
			op1BattleGroup: _op1,
			op2BattleGroup: _op2,
			winnerBattleGroup: 0,
			loserBattleGroup: 0
		});
		uint256 newID = battles.push(_battle) - 1;

		// Make sure we never overflow the battle max (4 billion battles)
		require(newID == uint256(uint32(newID)));

		// Determine Winner / Loser
		var (winnerID, loserID) = determineWinner(_battle);
		_battle.winnerBattleGroup = winnerID;
		_battle.loserBattleGroup = loserID;

		activateBattleCooldown(_battle); 	// Activate Battle Cooldown

		// Emit Battle Result Event
		BattleResult(
			newID,						// Battle ID
			_battle.winnerBattleGroup,	// Winner Battle Group ID
			_battle.loserBattleGroup	// Loser Battle Group ID
		);

		return newID;
	}

	function isReadyForBattle(uint256 _op) private view returns (bool) {
		return cryptoCardsContract.BattleGroupContract().isGroupReadyForBattle(_op);
	}

	function determineWinner(Battle _battle) private pure returns (uint256 winner, uint256 loser) {
		return (_battle.op2BattleGroup, _battle.op1BattleGroup);
	}

	function activateBattleCooldown(Battle _battle) private returns (bool) {
		cryptoCardsContract.BattleGroupContract().setGroupOnBattleCooldown(_battle.op1BattleGroup);
		cryptoCardsContract.BattleGroupContract().setGroupOnBattleCooldown(_battle.op2BattleGroup);
	}

	function countBattles() public view returns (uint) { // Return total count of all battles
		return battles.length;
	}
}