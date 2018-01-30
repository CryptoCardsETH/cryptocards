import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './pages/Home';
import Debug from './pages/Debug';
import Nav from './components/Nav';
import MetamaskInitialization from './components/MetamaskInitialization';

const App = () => (
  <Router>
    <div>
      <MetamaskInitialization />
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
