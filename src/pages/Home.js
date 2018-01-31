import React, { Component } from 'react';
import '../styles/App.css';
import { Jumbotron } from 'reactstrap';
import { APP_NAME } from '../config';
class HomePage extends Component {
  render() {
    return (
      <Jumbotron>
        <h1 className="display-3">{APP_NAME}</h1>
      </Jumbotron>
    );
  }
}

export default HomePage;
