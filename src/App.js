import React from 'react';
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  Switch,
  withRouter
} from 'react-router-dom';

import Home from './pages/Home';
import Debug from './pages/Debug';
import LoginPage from './pages/Login';
import Nav from './components/Nav';
import Web3Initialization from './components/Web3Initialization';
import { connect } from 'react-redux';

const PrivateRoute = ({
  component: Component,
  isAuthenticated,
  isAllowed,
  ...rest
}) => (
  <Route
    {...rest}
    render={props =>
      isAuthenticated ? (
        isAllowed ? (
          <Component {...props} />
        ) : (
          <div>
            <h1>Permission Denied</h1>You do not have the rights to access this
            page.
          </div>
        )
      ) : (
        <Redirect
          to={{
            pathname: '/login',
            state: { from: props.location }
          }}
        />
      )
    }
  />
);

const UserRoute = withRouter(
  connect(state => ({
    isAuthenticated: state.user.authenticated,
    isAllowed: true
  }))(PrivateRoute)
);

const App = () => (
  <Router>
    <div>
      <Web3Initialization />
      <Nav />
      <main role="main" className="container">
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/login" component={LoginPage} />
          <Route path="/debug" component={Debug} />

          {/*Routes that only logged in Users can access*/}
          <UserRoute path="/useronly" component={Debug} />
        </Switch>
      </main>
      {/*<Footer />*/}
    </div>
  </Router>
);
export default App;
