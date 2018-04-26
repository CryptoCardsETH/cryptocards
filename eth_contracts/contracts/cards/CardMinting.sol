pragma solidity ^0.4.17;
import "./CardBattles.sol";

contract CardMinting is CardBattles {
    function createGen0Card() external whenNotPaused returns (uint256 newCardID) {
        newCardID = createCard(0, 123456); // address msg.sender, attributes
    }
}
