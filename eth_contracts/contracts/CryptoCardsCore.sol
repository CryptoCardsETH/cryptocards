pragma solidity ^0.4.17;
import "./cards/CardMinting.sol";
import "./battles/Battles.sol";
import "./battles/BattleGroups.sol";
import "./battles/BattleQueue.sol";

contract CryptoCardsCore is CardMinting {
	Battles _battleContract;
	BattleGroups _battleGroupContract;
	BattleQueue _battleQueueContract;

	function CryptoCardsCore() public {
		ownerAddress = msg.sender;
		_battleContract = new Battles();
		_battleGroupContract = new BattleGroups();
		_battleQueueContract = new BattleQueue();
	}

	function BattleContract() public view returns (Battles) {
		return _battleContract;
	}

	function BattleGroupContract() public view returns (BattleGroups) {
		return _battleGroupContract;
	}

	function BattleQueueContract() public view returns (BattleQueue) {
		return _battleQueueContract;
	}

	function() external payable { // Only accept payments from our contracts
		require(
			msg.sender == address(_battleContract)
		);
	}
}