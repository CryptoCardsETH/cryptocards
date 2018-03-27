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
export const EDIT_ME_DETAILS = 'EDIT_ME_DETAILS';

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

//request to update email and nickname of user in account page
export function updateMe() {
  return (dispatch, getState) => {
    let meState = getState().user.me;
    dispatch(requestMe());
    return apiFetch('me', { method: 'PUT', body: JSON.stringify(meState) })
      .then(response => response.json())
      .then(json => {
        if (json.success) {
          toast.success('Changes Saved!');
          dispatch(receiveMe(json.data));
        } else toast.error('User not updated! ' + json.data);
        //todo: error checking (i.e. expired token)
      });
  };
}

//to edit the global state of user's info
export function editMeDetails(key, value) {
  return {
    type: EDIT_ME_DETAILS,
    key,
    value
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

export const REQUEST_USER_DETAIL = 'REQUEST_USER_DETAIL';
export const RECEIVE_USER_DETAIL = 'RECEIVE_USER_DETAIL';

export function fetchUserDetail(userId) {
  return dispatch => {
    dispatch(requestUserDetail());
    return apiFetch('users/' + userId)
      .then(response => response.json())
      .then(json => {
        dispatch(receiveUserDetail(userId, json.data));
      });
  };
}

function requestUserDetail(userId) {
  return {
    type: REQUEST_USER_DETAIL,
    userId
  };
}

function receiveUserDetail(userId, user) {
  return {
    type: RECEIVE_USER_DETAIL,
    userId,
    user
  };
}

export function follow(userId) {
  return dispatch => {
    return apiFetch('follow/' + userId, { method: 'PUT' })
      .then(response => response.json())
      .then(json => {
        if (json.success) {
          dispatch(fetchUserDetail(userId));
        } else toast.error('You are already a follower');
      });
  };
}
export const RECEIVE_ALL_USERS = 'RECEIVE_ALL_USERS';

export function fetchAllUsers() {
  return dispatch => {
    return apiFetch('users')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveAllUsers(json.data));
      });
  };
}

function receiveAllUsers(users) {
  return {
    type: RECEIVE_ALL_USERS,
    users
  };
}
