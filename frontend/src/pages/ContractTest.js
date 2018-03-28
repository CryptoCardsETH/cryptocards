import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Eth from 'ethjs-query';
import EthContract from 'ethjs-contract';
import { BattleGroupABI } from './../contracts';

class DebugPage extends Component {
  test = (senderAddress, contractAddress) => {
    const eth = new Eth(window.web3.currentProvider);
    const contract = new EthContract(eth);

    const miniToken = contract(BattleGroupABI).at(contractAddress);

    miniToken
      .createBattleGroup(senderAddress, [1, 2, 3, 4, 5], {
        from: senderAddress
      })
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
    let { loaded, addresses } = this.props.contract;
    let contractAddress = loaded ? addresses['BattleGroups'].address : null;
    return (
      <div>
        <button
          disabled={!loaded}
          onClick={() =>
            this.test(this.props.user.main_address, contractAddress)
          }
        >
          {loaded ? 'test contract' : 'loading...'}
        </button>
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user, contract } = state;
  return { user, contract };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({}, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(DebugPage);
