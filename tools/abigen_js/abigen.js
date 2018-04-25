const fs = require('fs');
const solc = require('solc');
const path = require('path');
const glob = require("glob");
const BASE = './eth_contracts/contracts/';
// https://github.com/ethereum/solc-js/issues/114#issuecomment-354752466
glob(BASE + "**/*.sol", function(er, files) {
    //read all the files in the contracts directory

    console.log(files);
    var inputs = {};
    files.forEach(f => {
        let contractName = path.basename(f).slice(0, -4);
        console.log("processing contract " + contractName);
        inputs[path.basename(f)] = fs.readFileSync(f).toString();
    });

    function findImports(path) {
        let aa = {
            "cards/CardMinting.sol": BASE+"cards/CardMinting.sol",
            "cards/CardOwnership.sol": BASE+"cards/CardOwnership.sol",
            "battles/Battles.sol": BASE+"battles/Battles.sol",
            "battles/BattleGroups.sol": BASE+"battles/BattleGroups.sol",
            "cards/CardBase.sol": BASE+"cards/CardBase.sol"
        };

        if(path in aa) {
            //hacky but referential import fix...
            path = aa[path];
        }

        return {
            'contents': fs.readFileSync(path).toString()
        };
    }
    // compile the solidity code
    var compiledCode = solc.compile({
        sources: inputs
    }, 1, findImports);

    // now, strip down each child to only have the contracts key
    let obj = compiledCode['contracts'];
    let data = {};
    Object.keys(obj).forEach(function(key) {
        data[key] = {
            'bytecode': obj[key].bytecode,
            'interface': JSON.parse(obj[key].interface)
        };
    });

    //build an exportable JS object, dump it into React's directory
    const fileContents =
          "//DO NOT EDIT; abigen.js generates this"
          + "\n const contracts = "
          + JSON.stringify(data, null, 2)
          + "; export default contracts;";
    fs.writeFile('/frontend/src/compiled_contracts.js', fileContents, function(err) {
        if (err) throw err;
        console.log('Compiled & saved');
    });

});
