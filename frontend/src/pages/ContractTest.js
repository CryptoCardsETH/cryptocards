import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { CONTRACT_NAME_BATTLEGROUPS } from './../contracts';
import {
  isReadyForContract,
  waitForTxToBeMined,
  getContractInstanceByName
} from '../selectors';
class ContractPlaygroundPage extends Component {
  test = (senderAddress, contractInstance) => {
    //todo: web3 check
    contractInstance
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
    let { ready, contractInstance } = this.props;
    return (
      <div>
        <button
          disabled={!ready}
          onClick={() =>
            this.test(this.props.user.main_address, contractInstance)
          }
        >
          {ready ? 'test contract' : 'loading...'}
        </button>
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user } = state;
  return {
    user,
    ready: isReadyForContract(state),
    contractInstance: getContractInstanceByName(
      state,
      CONTRACT_NAME_BATTLEGROUPS
    )
  };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({}, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(
  ContractPlaygroundPage
);
