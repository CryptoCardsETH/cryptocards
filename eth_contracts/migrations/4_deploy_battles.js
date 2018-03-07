var Battles = artifacts.require("./Battles");
var BattleGroups = artifacts.require("./BattleGroups");

module.exports = function(deployer) {
	deployer.deploy(Battles);
	deployer.deploy(BattleGroups);
};
