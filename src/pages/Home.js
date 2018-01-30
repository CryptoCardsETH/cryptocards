import React, { Component } from 'react';
import '../styles/App.css';
import { Jumbotron } from 'reactstrap';
class HomePage extends Component {
  render() {
    return (
      <div>
        <main role="main" className="container">
          <Jumbotron>
            <h1 className="display-3">cryptocards</h1>
          </Jumbotron>
        </main>
      </div>
    );
  }
}

export default HomePage;
