var BattleGroups = artifacts.require("BattleGroups");

module.exports = function(deployer) {
	deployer.deploy(BattleGroups);
};
