import React, { Component } from 'react';
import '../styles/App.scss';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Web3Login from '../components/Web3Login';
import { BooleanStatus } from '../components/Icons';
import { isReadyForContract } from '../selectors';
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
  render() {
    let { user, contract } = this.props;
    return (
      <div className="container">
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
              user,
              contract
            },
            true,
            2
          )}
        </pre>
        <Web3Login />
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user, contract } = state;
  return { user, contract, a: isReadyForContract(state) };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({}, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(DebugPage);
