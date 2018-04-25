pragma solidity ^0.4.17;
import "./CardOwnership.sol";

contract CardBattles is CardOwnership {
	uint64 public constant BATTLE_COOLDOWN_TIME = 60*60;

	function isReadyForBattle(uint256 _cardID) public view returns (bool) {
		requireCardExists(_cardID);
		if (cards[_cardID].battleCooldownEnd > uint64(now)) {
			return false;
		}
		return true;
	}

	function setOnBattleCooldown(uint256 _cardID) public returns (bool) {
		requireCardExists(_cardID);
		cards[_cardID].battleCooldownEnd = uint64(now) + BATTLE_COOLDOWN_TIME;
		return true;
	}
}
