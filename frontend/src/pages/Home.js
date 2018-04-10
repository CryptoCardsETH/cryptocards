import React, { Component } from 'react';
import '../styles/App.scss';
import { Jumbotron } from 'reactstrap';
import { APP_NAME } from '../config';
import Web3Login from '../components/Web3Login';
class HomePage extends Component {
  render() {
    return (
      <div className="container home">
        <div className="row">
          <div className="col-md-6">
            {/* <h1 className="text-center">{APP_NAME}</h1> */}
            <h1 className="text-center opening">Collect</h1>
            <h1 className="text-center opening">Trade</h1>
            <h1 className="text-center opening">Battle</h1>
            <br />
            {/* <p className="text-center">Collect, <br/> Trade, <br/> Battle <br/></p> */}
            <div className="login text-center">
              <Web3Login />
            </div>
          </div>
          <div className="col-md-6">
            <img src="http://files.harrischristiansen.com/0L1R2p3G2s05/cryptocards_logo.png" />
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
