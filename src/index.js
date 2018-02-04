import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import 'bootstrap/dist/css/bootstrap.css';
import App from './App';
import configureStore from './store/configureStore';
import { Provider } from 'react-redux';
import registerServiceWorker from './registerServiceWorker';
import { getLocally } from './actions/index';
import { loginFromJWT } from './actions/users';

const store = configureStore();

const jwt = getLocally('jwt');
if (jwt) store.dispatch(loginFromJWT(jwt));

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
registerServiceWorker();
