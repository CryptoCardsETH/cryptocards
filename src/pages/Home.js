import React, { Component } from 'react';
import '../styles/App.scss';
import { Jumbotron } from 'reactstrap';
import { APP_NAME } from '../config';
import Web3Login from '../components/Web3Login';
class HomePage extends Component {
  render() {
    return (
      <div>
        <Jumbotron>
          <h1 className="display-3">{APP_NAME}</h1>
        </Jumbotron>
        <Web3Login />
      </div>
    );
  }
}

export default HomePage;
