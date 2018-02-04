import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './pages/Home';
import Debug from './pages/Debug';
import Nav from './components/Nav';
import Web3Initialization from './components/Web3Initialization';

const App = () => (
  <Router>
    <div>
      <Web3Initialization />
      <Nav />
      <main role="main" className="container">
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/debug" component={Debug} />
        </Switch>
      </main>
      {/*<Footer />*/}
    </div>
  </Router>
);
export default App;
