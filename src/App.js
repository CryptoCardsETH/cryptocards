import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './pages/Home';
import { Container } from 'reactstrap';

const App = () => (
  <Router>
    <div>
      {/*<Nav />*/}
      <Container>
        <Switch>
          <Route exact path="/" component={Home} />
        </Switch>
      </Container>
      {/*<Footer />*/}
    </div>
  </Router>
);
export default App;
