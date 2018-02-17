import apiFetch from './index';

export const REQUEST_ALL_CARDS = 'REQUEST_ALL_CARDS';
export const RECEIVE_ALL_CARDS = 'RECEIVE_ALL_CARDS';

export function fetchAllCards() {
  return dispatch => {
    dispatch(requestAllCards());
    return apiFetch('cards')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveAllCards(json.data));
      });
  };
}

function requestAllCards() {
  return {
    type: REQUEST_ALL_CARDS
  };
}

function receiveAllCards(cards) {
  return {
    type: RECEIVE_ALL_CARDS,
    cards
  };
}
