import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import 'bootstrap/dist/css/bootstrap.css';
import App from './App';
import configureStore from './store/configureStore';
import { Provider } from 'react-redux';
import registerServiceWorker from './registerServiceWorker';
import { getLocally } from './actions/index';
import { loginFromJWT, setSignedMessagesBulk } from './actions/users';

const store = configureStore();

const jwt = getLocally('jwt');
if (jwt) store.dispatch(loginFromJWT(jwt));
const signedMessages = getLocally('signedMessages');
if (signedMessages) store.dispatch(setSignedMessagesBulk(signedMessages));
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
registerServiceWorker();
