var CryptoCardsCore = artifacts.require("./CryptoCardsCore");

module.exports = function(deployer) {
	return deployer.deploy(CryptoCardsCore);
};
