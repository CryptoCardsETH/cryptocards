import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './pages/Home';
import Nav from './components/Nav';

const App = () => (
  <Router>
    <div>
      <Nav />
      <main role="main" className="container">
        <Switch>
          <Route exact path="/" component={Home} />
        </Switch>
      </main>
      {/*<Footer />*/}
    </div>
  </Router>
);
export default App;
