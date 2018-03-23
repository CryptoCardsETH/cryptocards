var CardOwnership = artifacts.require("./CardOwnership");

module.exports = function(deployer) {
	// Deploy CardBase, then CardOwnership
	return deployer.deploy(CardOwnership);
};
