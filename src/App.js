import React, { Component } from 'react';
import './App.css';
import { Jumbotron } from 'reactstrap';
class App extends Component {
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

export default App;
