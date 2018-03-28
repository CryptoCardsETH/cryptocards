import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Web3Login from '../components/Web3Login';
import { BooleanStatus } from '../components/Icons';
import Eth from 'ethjs-query';
import EthContract from 'ethjs-contract';
import CompiledContracts from './../compiled_contracts';

class DebugPage extends Component {
  //https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md#list-of-chain-ids
  getCanonicalNetworkName = networkId => {
    switch (networkId) {
      case 1:
        return 'mainnet (1)';
      case 2:
        return 'deprecated Morden test network (2)';
      case 3:
        return 'ropsten test network (3)';
      case 4:
        return 'rinkeby test network (4)';
      case 42:
        return 'kovan test network (42)';
      default:
        return 'unknown network (' + networkId + ')';
    }
  };
  test = () => {
    const senderAddress = '0xAB90357976081f106B785910237a9D1fA9b05ED8';
    const address = '0x4b591a1068e915cee06eb0da97e653ecce8a8777';
    const contractData = CompiledContracts['BattleGroups.sol:BattleGroups'];
    const eth = new Eth(window.web3.currentProvider);
    const contract = new EthContract(eth);

    const MiniToken = contract(contractData['interface']);
    const miniToken = MiniToken.at(address);

    miniToken
      .createBattleGroup(
        '0xAB90357976081f106B785910237a9D1fA9b05ED8',
        [4, 3, 4, 2, 3],
        { from: senderAddress }
      )
      .then(txHash => {
        console.log('Transaction sent');
        console.dir(txHash);
        this.waitForTxToBeMined(txHash, eth);
      })
      .catch(console.error);
  };

  async waitForTxToBeMined(txHash, eth) {
    let txReceipt;
    while (!txReceipt) {
      try {
        txReceipt = await eth.getTransactionReceipt(txHash);
      } catch (err) {
        return console.log(err);
      }
    }
    console.log('yayyayayay');
  }

  render() {
    let { user } = this.props;
    return (
      <div>
        <h1>Debug</h1>
        <BooleanStatus bool={user.isWeb3Available} /> web3 available
        <br />
        <BooleanStatus bool={user.accounts_list.length >= 1} /> at least one eth
        account
        <br />
        <BooleanStatus bool={user.authenticated} /> authenticated
        <br />
        <b>network: </b>
        {this.getCanonicalNetworkName(user.network_id)}
        <pre>
          {JSON.stringify(
            {
              user
            },
            true,
            2
          )}
        </pre>
        <Web3Login />
        <button onClick={this.test}>testcontract</button>
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user } = state;
  return { user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({}, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(DebugPage);
