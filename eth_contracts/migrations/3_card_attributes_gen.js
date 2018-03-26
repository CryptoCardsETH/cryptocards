var GenerateCard = artifacts.require("./GenerateCard");

module.exports = function(deployer) {
	// GenerateCard
	deployer.deploy(GenerateCard);
};
