import apiFetch, { persistLocally } from './index';
import jwtDecode from 'jwt-decode';
import { toast } from 'react-toastify';
export const LOGIN_FROM_JWT_SUCCESS = 'LOGIN_FROM_JWT_SUCCESS';
export function loginFromJWT(token) {
  return dispatch => {
    let decoded = jwtDecode(token);

    let expirationDate = decoded.exp;
    let timeMs = Math.round(new Date().getTime() / 1000);
    console.log('DECODED', decoded);

    let shouldFetch = true;
    if (expirationDate) {
      if (expirationDate >= timeMs) console.log('expiring soon!');
      else if (expirationDate < timeMs) {
        console.log('expired!!');
        toast.error('Token expired! Please log in again');
        shouldFetch = false;
        dispatch(logout());
      }
    }
    if (shouldFetch) {
      persistLocally('jwt', token);
      dispatch(saveToken(token));
      setTimeout(() => {
        dispatch(fetchMe());
      }, 50);
    }
  };
}

function saveToken(token) {
  return {
    type: LOGIN_FROM_JWT_SUCCESS,
    token: token
  };
}

export const REMOVE_TOKEN = 'REMOVE_TOKEN';
function removeToken() {
  return {
    type: REMOVE_TOKEN
  };
}

export function logout() {
  return dispatch => {
    persistLocally('jwt', null);
    dispatch(removeToken());
  };
}

export const REQUEST_ME = 'REQUEST_ME';
export const RECEIVE_ME = 'RECEIVE_ME';

export function fetchMe() {
  return dispatch => {
    dispatch(requestMe());
    return apiFetch('me')
      .then(response => response.json())
      .then(json => {
        //todo: error checking (i.e. expired token)
        dispatch(receiveMe(json.data));
      });
  };
}

function requestMe() {
  return {
    type: REQUEST_ME
  };
}

function receiveMe(json) {
  return {
    type: RECEIVE_ME,
    me: json
  };
}

export const SET_WEB3_AVAILABILITY = 'SET_WEB3_AVAILABILITY';
export function setWeb3Availability(isAvailable) {
  return { type: SET_WEB3_AVAILABILITY, isAvailable };
}
export const SET_SIGNED_MESSAGES = 'SET_SIGNED_MESSAGES';
function setSignedMessage(address, message) {
  return (dispatch, getState) => {
    let { signedMessages } = getState().user;
    signedMessages[address] = message;
    persistLocally('signedMessages', signedMessages);
    return { type: SET_SIGNED_MESSAGES, signedMessages };
  };
}
export function setSignedMessagesBulk(signedMessages) {
  return { type: SET_SIGNED_MESSAGES, signedMessages };
}
export const SET_NETWORK_ID = 'SET_NETWORK_ID';
export function setNetworkId(network_id) {
  return { type: SET_NETWORK_ID, network_id };
}

export const SET_ACCOUNTS_LIST = 'SET_ACCOUNTS_LIST';
export function setAccountsList(accounts) {
  return { type: SET_ACCOUNTS_LIST, accounts };
}

export function initializeAuthFlow(address, signed) {
  return dispatch => {
    dispatch(setSignedMessage(address, signed));
    return apiFetch('auth', {
      method: 'POST',
      body: JSON.stringify({ address, signed, plaintext: 'CryptoCards' })
    })
      .then(response => response.json())
      .then(json => {
        if (json.success) {
          dispatch(loginFromJWT(json.data.token));
        } else {
          console.log('oops');
        }
      });
  };
}

export const REQUEST_MY_CARDS = 'REQUEST_MY_CARDS';
export const RECEIVE_MY_CARDS = 'RECEIVE_MY_CARDS';

export function fetchMyCards() {
  return dispatch => {
    dispatch(requestMyCards());
    return apiFetch('me/cards')
      .then(response => response.json())
      .then(json => {
        //todo: error checking (i.e. expired token)
        dispatch(receiveMyCards(json.data));
      });
  };
}

function requestMyCards() {
  return {
    type: REQUEST_MY_CARDS
  };
}

function receiveMyCards(cards) {
  return {
    type: RECEIVE_MY_CARDS,
    cards
  };
}
