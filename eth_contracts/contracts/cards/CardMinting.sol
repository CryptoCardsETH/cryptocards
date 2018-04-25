pragma solidity ^0.4.17;
import "./CardBattles.sol";

contract CardMinting is CardBattles {
    function createGen0Card() external returns (uint256 newCardID) {
        uint256 _newCardID = createCard(0, 123456); // address msg.sender, attributes
        return _newCardID;
    }
}
