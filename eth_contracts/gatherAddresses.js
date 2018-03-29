const fs = require('fs');
const path = require('path');
const BASE = 'build/contracts/';
const NETWORK_ID = 999999;

let addresses = {};
fs.readdirSync(BASE).forEach(f=> {
    let c = JSON.parse(fs.readFileSync(BASE+f).toString());
    let { contractName, networks } = c;
    console.log(`Processing Build atrifact for contract ${contractName} from ${f}`);
    let targetNetwork = networks[NETWORK_ID];
    if(targetNetwork){
        let { address, transactionHash } = targetNetwork;
        addresses[contractName] = {address, transactionHash};
    }
});

console.log(addresses);


var http = require("http");

var options = {
    "method": "PUT",
    "hostname": "nginx",
    "port": "8080",
    "path": "/v1/contracts/ingest",
    "headers": {
        "content-type": "application/json"
    }
};

var req = http.request(options, function (res) {
    var chunks = [];

    res.on("data", function (chunk) {
        chunks.push(chunk);
    });

    res.on("end", function () {
        var body = Buffer.concat(chunks);
        console.log(body.toString());
    });
});

req.write(JSON.stringify(addresses));
req.end();
