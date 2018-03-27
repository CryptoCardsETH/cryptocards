const fs = require('fs');
const solc = require('solc');
const path = require('path');
const glob = require("glob");
const BASE = './eth_contracts/contracts/';
// https://github.com/ethereum/solc-js/issues/114#issuecomment-354752466
glob(BASE + "**/*.sol", function(er, files) {
    console.log(files);
    var inputs = {};
    files.forEach(f => {
        let contractName = path.basename(f).slice(0, -4);
        console.log("processing contract " + contractName);
        inputs[path.basename(f)] = fs.readFileSync(f).toString();
    });

    function findImports(path) {
        if(path == "contracts/CardBase.sol") {
            //hacky but...
            path = BASE+"CardBase.sol";
        }
        return {
            'contents': fs.readFileSync(path).toString()
        };
    }

    var compiledCode = solc.compile({
        sources: inputs
    }, 1, findImports);
    console.log(compiledCode);

    let obj = compiledCode['contracts'];
    let data = {};
    Object.keys(obj).forEach(function(key) {
        data[key] = {
            'bytecode': obj[key].bytecode,
            'interface': JSON.parse(obj[key].interface)
        };
    });

    fs.writeFile('/frontend/src/compiled_contracts.js', "export default contracts = " + JSON.stringify(data, null, 2), function(err) {
        if (err) throw err;
        console.log('Compiled & saved');
    });

});
