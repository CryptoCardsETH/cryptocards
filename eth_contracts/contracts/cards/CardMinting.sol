pragma solidity ^0.4.17;
import "./CardOwnership.sol";

contract CardMinting is CardOwnership {
    function createGen0Card() external returns (uint256 isSuccess) {
        createCard(0, 123456); // address msg.sender, attributes
        return 1;
    }
}
