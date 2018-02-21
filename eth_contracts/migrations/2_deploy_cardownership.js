var CardBase = artifacts.require("CardBase");
var CardOwnership = artifacts.require("CardOwnership");

module.exports = function(deployer) {
	// Deploy CardBase, then CardOwnership
	deployer.deploy(CardBase).then(function() {
		return deployer.deploy(CardOwnership, CardBase.address);
	});
};
