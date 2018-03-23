module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    development: {
      host: process.env.RPC_HOST || "127.0.0.1",
      port: process.env.RPC_PORT || 7545,
      network_id: "*" // Match any network id
      //gas: 4600000
    }
  }
};
