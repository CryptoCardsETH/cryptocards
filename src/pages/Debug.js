import React, { Component } from 'react';
import '../styles/App.css';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import MetamaskLogin from '../components/MetamaskLogin';
import { BooleanStatus } from '../components/Icons';

class DebugPage extends Component {
  render() {
    let { user } = this.props;
    return (
      <div>
        <h1>Debug</h1>
        <BooleanStatus bool={user.isWeb3Available} /> web3 available
        <br />
        <BooleanStatus
          bool={user.isWeb3Available && window.web3.eth.accounts.length >= 1}
        />{' '}
        at least one eth account
        <br />
        <BooleanStatus bool={user.authenticated} /> authenticated
        <pre>
          {JSON.stringify(
            {
              accounts: user.isWeb3Available && window.web3.eth.accounts,
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
