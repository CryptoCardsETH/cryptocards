import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Web3Login from '../components/Web3Login';
import { BooleanStatus } from '../components/Icons';
import Eth from 'ethjs-query';
import EthContract from 'ethjs-contract';
import { BattleGroup } from './../contracts';

class DebugPage extends Component {
  test = senderAddress => {
    const eth = new Eth(window.web3.currentProvider);
    const contract = new EthContract(eth);

    const miniToken = contract(BattleGroup.ABI).at(BattleGroup.address);

    miniToken
      .createBattleGroup(senderAddress, [4, 3, 4, 2, 3], {
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
    let { user } = this.props;
    return (
      <div>
        <button onClick={() => this.test(user.main_address)}>
          testcontract
        </button>
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
