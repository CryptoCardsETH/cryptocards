import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { initializeAuthFlow } from '../actions/users';

import { Button } from 'reactstrap';

import ethUtil from 'ethereumjs-util';
import sigUtil from 'eth-sig-util';
import Web3 from 'web3';
import { APP_NAME } from '../config';
import { toast } from 'react-toastify';

class Web3Login extends React.Component {
  //sign the message containing a simple string, which will we give to server in exchange for a JWT
  signMessage = (fromAddress, web3client) => {
    let hexEncodedMessage = ethUtil.bufferToHex(new Buffer(APP_NAME, 'utf8'));
    web3client.currentProvider.sendAsync(
      {
        method: 'personal_sign',
        params: [hexEncodedMessage, fromAddress],
        fromAddress
      },
      (err, result) => {
        if (err) return console.error(err);
        if (result.error) return console.error(result.error);
        console.log(`PERSONAL SIGNED: ${result.result}`);
        this.props.initializeAuthFlow(fromAddress, result.result);

        //test code to make sure that it was signed right. this is what the backend'll do
        console.log('recovering...');
        const msgParams = { data: hexEncodedMessage };
        msgParams.sig = result.result;
        console.dir({ msgParams });
        const recovered = sigUtil.recoverPersonalSignature(msgParams);
        console.dir({ recovered });
        if (recovered === fromAddress) {
          console.log(`SigUtil Successfully verified signer as ${fromAddress}`);
        } else {
          console.dir(recovered);
          console.log(
            `SigUtil Failed to verify signer when comparing ${
              recovered.result
            } to ${fromAddress}`
          );
          console.log('Failed, comparing %s to %s', recovered, fromAddress);
        }
      }
    );
  };
  doLogin = () => {
    if (window.web3) {
      const web3client = new Web3(window.web3.currentProvider);
      let { accounts_list } = this.props.user;
      if (accounts_list.length === 0) {
        console.log('no accounts!');
        toast.error('No ETH accounts! Perhaps metamask is locked?');
      } else {
        //todo: offer mechanism for picking address if there are multiple?
        this.signMessage(accounts_list[0], web3client);
      }
    } else {
      console.log('no web3');
      toast.error('No web3 client!');
    }
  };
  render() {
    return <Button onClick={this.doLogin}>Login</Button>;
  }
}

function mapStateToProps(state) {
  let { user } = state;
  return { user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      initializeAuthFlow
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(Web3Login);
