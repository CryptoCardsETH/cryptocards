pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CardBase.sol";

contract TestCardOwnership {
  CardOwnership cardOwnership = CardOwnership(DeployedAddresses.CardOwnership());

	// Testing the createCard() and ownerOf() function
	function testCreateCard() public {
		uint expected = cardOwnership.totalSupply();

		uint returnedID = cardOwnership.createCard(this, 0);

		Assert.equal(returnedID, expected, "Created card should have next sequential cardID");

		address cardOwner = cardOwnership.ownerOf(returnedID);

		Assert.equal(cardOwner, this, "Created card should be owned by this address");
	}

	// Testing the balanceOf() function
	function testBalanceOf() public {
		uint256 balance = cardOwnership.balanceOf(this);

		Assert.equal(balance, 1, "Test should own one token via balanceOf");
	}

	// Test creating many additional cards
	function testCreateManyCards() public {
		uint numCards = 5;
		uint expected = cardOwnership.balanceOf(this) + numCards;

		uint i = 0;
		for (i = 0; i < numCards; i++) {
			cardOwnership.createCard(this);
		}

		uint256 balance = cardOwnership.balanceOf(this);
		Assert.equal(balance, expected, "Balance should increase to expected count after creating many cards");
	}

	// Testing the tokensOfOwner() function
	/*function testTokensOfOwner() public {
		uint256[] memoryownerTokens = cardOwnership.tokensOfOwner(this);

		Assert.equal(ownerTokens.length, 1, "Test should own one token via tokensOfOwner");
	}*/

}
