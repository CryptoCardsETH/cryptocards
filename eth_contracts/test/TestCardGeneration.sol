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
            uint256 isSuccess = cryptoCards.createGen0Card();
            Assert.equal(isSuccess, uint256(1), "Generation of gen 0 card should succeed");
        }

        uint256 totalCount = cryptoCards.totalSupply();
        Assert.equal(totalCount, expected, "Total count after card generation is wrong");

    }
}
