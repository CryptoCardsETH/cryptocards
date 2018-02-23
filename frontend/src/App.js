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
import AllCards from './pages/AllCards';
import LoginPage from './pages/Login';
import MarketplacePage from './pages/Marketplace';
import AccountPage from './pages/Account';
import TransactionPage from './pages/AllTransactions';
import Nav from './components/Nav';
import Web3Initialization from './components/Web3Initialization';
import { connect } from 'react-redux';
import { ToastContainer } from 'react-toastify';
import CardDetail from './pages/CardDetail';
import FAQ from './pages/Static/FAQ';
import UserDetail from './pages/UserDetail';

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
      <ToastContainer />
      <Nav />
      <main role="main" className="container" style={{ paddingTop: '80px' }}>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/faq" component={FAQ} />

          <Route path="/login" component={LoginPage} />
          <Route path="/debug" component={Debug} />
          <Route path="/cards" component={AllCards} />
          <Route path="/card/:id" component={CardDetail} />
          <Route path="/user/:id" component={UserDetail} />
          <Route path="/marketplace" component={MarketplacePage} />

          {/*Routes that only logged in Users can access*/}
          <UserRoute path="/useronly" component={Debug} />
          <UserRoute path="/account" component={AccountPage} />
          <UserRoute path="/transactions" component={TransactionPage} />
        </Switch>
      </main>
      {/*<Footer />*/}
    </div>
  </Router>
);
export default App;
