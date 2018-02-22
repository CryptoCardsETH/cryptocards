pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CardBase.sol";
import "../contracts/GenerateCard.sol";

contract TestCardGeneration {

    CardOwnership cardOwnership = CardOwnership(DeployedAddresses.CardOwnership());

    GenerateCard generateCard = GenerateCard(DeployedAddresses.GenerateCard());



    // Testing the createCard() function
    function testCreateCard() public {
        uint256 numCards = 50;
        uint256 expected = cardOwnership.totalSupply() + numCards;

        generateCard.setCardOwnership(cardOwnership);
        uint256 returnVal = generateCard.initialGeneration(numCards);
        Assert.equal(returnVal, uint256(1), "Generation of cards should succeed");

        uint256 totalCount = cardOwnership.totalSupply();
        Assert.equal(totalCount, expected, "Total count after card generation is wrong");

    }
}
