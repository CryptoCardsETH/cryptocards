import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { setWeb3Availability } from '../actions/users';

class MetamaskInitialization extends React.Component {
  componentDidMount() {
    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (window.web3) {
      // Use the browser's ethereum provider
      let provider = window.web3.currentProvider;
      this.props.setWeb3Availability(true);
      console.log(provider);
    } else {
      console.log('No web3? You should consider trying MetaMask!');
      this.props.setWeb3Availability(false);
    }
  }

  render() {
    return null;
  }
}

function mapStateToProps(state) {
  let { user } = state;
  return { user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      setWeb3Availability
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(
  MetamaskInitialization
);
