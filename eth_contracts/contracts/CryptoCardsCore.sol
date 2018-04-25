pragma solidity ^0.4.17;
import "./cards/CardMinting.sol";
import "./battles/Battles.sol";
import "./battles/BattleGroups.sol";

contract CryptoCardsCore is CardMinting {
    address public ownerAddress;
    Battles _battleContract;
    BattleGroups _battleGroupContract;
    function CryptoCardsCore() public {
        ownerAddress = msg.sender;
        _battleContract = new Battles();
        _battleGroupContract = new BattleGroups();
    }

    function BattleContract() public view returns (Battles) {
        return _battleContract;
    }

    function BattleGroupContract() public view returns (BattleGroups) {
        return _battleGroupContract;
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