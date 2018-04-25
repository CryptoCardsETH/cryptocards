pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CryptoCardsCore.sol";

contract TestCardOwnership {
  CryptoCardsCore cryptoCards = CryptoCardsCore(DeployedAddresses.CryptoCardsCore());

	// Testing the createCard() and ownerOf() function
	function testCreateCard() public {
		uint expected = cryptoCards.totalSupply();

		uint returnedID = cryptoCards.createCard(this, 0);

		Assert.equal(returnedID, expected, "Created card should have next sequential cardID");

		address cardOwner = cryptoCards.ownerOf(returnedID);

		Assert.equal(cardOwner, this, "Created card should be owned by this address");
	}

	// Testing the balanceOf() function
	function testBalanceOf() public {
		uint256 balance = cryptoCards.balanceOf(this);

		testCreateCard();
		Assert.equal(balance+1, cryptoCards.balanceOf(this), "Result of balanceOf changed by an unexpected amount.");

		testCreateCard();
		Assert.equal(balance+2, cryptoCards.balanceOf(this), "Result of balanceOf changed by an unexpected amount.");
	}

	// Test creating many additional cards
	function testCreateManyCards() public {
		uint numCards = 5;
		uint expected = cryptoCards.balanceOf(this) + numCards;

		uint i = 0;
		for (i = 0; i < numCards; i++) {
			testCreateCard();
		}

		uint256 balance = cryptoCards.balanceOf(this);
		Assert.equal(balance, expected, "Balance should increase to expected count after creating many cards");
	}

	// Testing the tokensOfOwner() function
	/*function testTokensOfOwner() public {
		uint256[] memoryownerTokens = cryptoCards.tokensOfOwner(this);

		Assert.equal(ownerTokens.length, 1, "Test should own one token via tokensOfOwner");
	}*/

}
