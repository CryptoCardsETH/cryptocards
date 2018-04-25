pragma solidity ^0.4.17;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CryptoCardsCore.sol";

contract TestCardGeneration {
    CryptoCardsCore cryptoCards = CryptoCardsCore(DeployedAddresses.CryptoCardsCore());

    function testGen0Creation() public {
        uint256 numCards = 50;
        uint256 expected = cryptoCards.totalSupply() + numCards;

        uint i = 0;
        for (i = 0; i < numCards; i++) {
            cryptoCards.createGen0Card();
        }

        uint256 totalCount = cryptoCards.totalSupply();
        Assert.equal(totalCount, expected, "Total count after card generation is wrong");

    }
}
