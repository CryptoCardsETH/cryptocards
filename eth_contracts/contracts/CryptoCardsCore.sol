pragma solidity ^0.4.17;
import "./cards/CardMinting.sol";
import "./battles/Battles.sol";
import "./battles/BattleGroups.sol";
import "./battles/BattleQueue.sol";

contract CryptoCardsCore is CardMinting {
	address _ownerAddress;
	Battles _battleContract;
	BattleGroups _battleGroupContract;
	BattleQueue _battleQueueContract;

	function CryptoCardsCore() public {
		_ownerAddress = msg.sender;
		_battleContract = new Battles();
		_battleGroupContract = new BattleGroups();
		_battleQueueContract = new BattleQueue();

		// Create Imaginary (Reference) Entries
		uint256 c = createCard(0, 34433);
		_battleGroupContract.createBattleGroup(this, [c,c,c,c,c]);
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

	address public newContractAddress;
	event ContractUpgrade(address newContract);
	function setNewAddress(address _newAddress) external {
		newContractAddress = _newAddress;
		ContractUpgrade(_newAddress); // Emit event announcing update to new contract
	}

	function() external payable { // Only accept payments from our contracts
		require(
			msg.sender == address(_battleContract)
		);
	}
}