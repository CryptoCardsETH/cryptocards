pragma solidity ^0.4.17;

import "../contracts/CardBase.sol";

// Generation 0 Card Creation
contract GenerateCard {

    CardOwnership public cardOwnership = new CardOwnership();

    function initialGeneration(uint256 numberOfCards) external returns (uint256 isSuccess) {
        uint256 i;
        for (i = uint256(1); i <= numberOfCards; i++) {
            uint256 returnedID = 1;
            cardOwnership.createCard(this, i);
            //cardOwnership.totalSupply();
            if (returnedID == 0) {
                return 0;
            }
        }
        return 1;
    }

    function setCardOwnership(address _cardOwnership) external {
        cardOwnership = CardOwnership(_cardOwnership);
    }

}
