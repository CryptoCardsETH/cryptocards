import React, { Component } from 'react';
import '../styles/App.css';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import MetamaskLogin from '../components/MetamaskLogin';
class DebugPage extends Component {
  render() {
    return (
      <div>
        <h1>Debug</h1>

        <pre>
          {JSON.stringify(
            { accounts: window.web3.eth.accounts, user: this.props.user },
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
