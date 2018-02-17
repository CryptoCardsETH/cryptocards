import React, { Component } from 'react';
import '../styles/App.css';
import Web3Login from '../components/Web3Login';
class HomePage extends Component {
  render() {
    return (
      <div>
        <h1 className="display-3">You need to login!</h1>
        <Web3Login />
      </div>
    );
  }
}

export default HomePage;
