import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { BattleGroupABI } from './../contracts';
import {
  getContractInstance,
  waitForTxToBeMined
} from './../actions/contracts';
class ContractPlaygroundPage extends Component {
  test = (senderAddress, contractAddress) => {
    //todo: web3 check
    const miniToken = getContractInstance(BattleGroupABI, contractAddress);
    miniToken
      .createBattleGroup(senderAddress, [1, 2, 3, 4, 5], {
        from: senderAddress
      })
      .then(txHash => {
        console.log('Transaction sent');
        console.dir(txHash);
        waitForTxToBeMined(txHash);
      })
      .catch(console.error);
  };

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
export default connect(mapStateToProps, mapDispatchToProps)(
  ContractPlaygroundPage
);
