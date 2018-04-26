pragma solidity ^0.4.17;
import "../CryptoCardsCore.sol";

contract BattleQueue {
	
	uint256 queue_battleGroupID; // ID of Battle Group in Queue

	// Queue Joined Event: Emitted every time a BattleGroup joins the Queue
	event QueueJoined(uint256 battleGroupID, uint64 eventTime);

	// Queue Matched Event: Emitted every time a match is made by the Queue
	event QueueMatched(uint256 battleGroup1, uint256 battleGroup2, uint64 eventTime);

	CryptoCardsCore public cryptoCardsContract; // Pointer to Card Contract
	function BattleQueue() public {
		CryptoCardsCore candidateContract = CryptoCardsCore(msg.sender);
		//require(candidateContract.supportsInterface(InterfaceSignature_ERC721));
		cryptoCardsContract = candidateContract;
	}

	function joinQueue(uint256 _battleGroupID) public returns (bool) {
		// Ensure Ready For Battle
		if (!isReadyForBattle(_battleGroupID)) {
			return false;
		}

		// Emit Queue Joined Event
		QueueJoined(
			_battleGroupID,		// Battle Group ID
			uint64(now)			// Event Time
		);

		// Add to Queue if no existing opponent
		if (queue_battleGroupID == 0) {
			queue_battleGroupID = _battleGroupID;
			return true;
		}

		// Emit Queue Matched Event
		QueueMatched(
			_battleGroupID,			// Battle Group 1 ID
			queue_battleGroupID,	// Battle Group 2 ID
			uint64(now)				// Event Time
		);

		// Create Battle
		sendToBattle(_battleGroupID, queue_battleGroupID);

		
		// Clear Queue of existing opponent
		queue_battleGroupID = 0;

		return true;
	}

	function isReadyForBattle(uint256 _op) private view returns (bool) {
		return cryptoCardsContract.BattleGroupContract().isGroupReadyForBattle(_op);
	}

	function sendToBattle(uint256 _op1, uint256 _op2) private returns (uint256) {
		return cryptoCardsContract.BattleContract().createBattle(_op1, _op2);
	}
}