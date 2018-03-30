var Blacklist = artifacts.require('./Blacklist'); 

module.exports = function(deployer) {
  deployer.deploy(Blacklist);  
}
