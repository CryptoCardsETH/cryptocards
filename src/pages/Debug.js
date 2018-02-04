import React, { Component } from 'react';
import '../styles/App.css';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import MetamaskLogin from '../components/MetamaskLogin';
import Web3 from 'web3';
import { BooleanStatus } from '../components/Icons';

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
    let { user } = this.props;
    let web3client = null;
    // if (window.web3) {
    //   web3client = new Web3(window.web3.currentProvider);
    // }
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
        <b>network:</b>
        {this.getCanonicalNetworkName(user.network_id)}
        <pre>
          {JSON.stringify(
            {
              accounts: user.accounts_list,
              user
            },
            true,
            2
          )}
        </pre>
        <MetamaskLogin />
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
