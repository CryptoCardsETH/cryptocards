import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import {
  setAccountsList,
  setNetworkId,
  setWeb3Availability
} from '../actions/users';
import Web3 from 'web3';
class Web3Initialization extends React.Component {
  componentDidMount() {
    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (window.web3 !== 'undefined') {
      this.props.setWeb3Availability(true);
      let web3client = new Web3(window.web3.currentProvider);
      web3client.eth.net.getId().then(id => {
        console.log('network id:', id);
        this.props.setNetworkId(id);
      });
      web3client.eth.getAccounts().then(list => {
        this.props.setAccountsList(list);
      });
    } else {
      console.log('No web3!');
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
      setWeb3Availability,
      setNetworkId,
      setAccountsList
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(Web3Initialization);
