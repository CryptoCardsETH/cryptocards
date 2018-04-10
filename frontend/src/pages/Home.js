import React, { Component } from 'react';
import '../styles/App.scss';
import { Jumbotron } from 'reactstrap';
import { APP_NAME } from '../config';
import Web3Login from '../components/Web3Login';
import 'animate.css';
class HomePage extends Component {
  render() {
    return (
      <div className="container-fluid home">
        <div className="row intro">
          <div className="col-md-12 intro-text">
            <h1> Welcome to {APP_NAME} </h1>
            <p>Collect, Trade, Battle</p>
            <Web3Login />
          </div>
          <div
            className="col-md-12 intro-image animated fadeInUp"
            style={{ animationDelay: '0.5s' }}
          >
            <img src="http://files.harrischristiansen.com/0L1R2p3G2s05/cryptocards_logo.png" />
          </div>
        </div>
        <div className="row secondary">
          <div className="col-md-12">
            <h1 className="text-center">Ownership Backed By the Blockchain</h1>
            <p className="text-center">Lorem Ipsum</p>
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;
