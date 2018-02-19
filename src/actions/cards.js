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
export const REQUEST_CARD_DETAIL = 'REQUEST_CARD_DETAIL';
export const RECEIVE_CARD_DETAIL = 'RECEIVE_CARD_DETAIL';

export function fetchCardDetail(cardId) {
  return dispatch => {
    dispatch(requestCardDetail());
    return apiFetch('cards/' + cardId)
      .then(response => response.json())
      .then(json => {
        dispatch(receiveCardDetail(cardId, json.data));
      });
  };
}

function requestCardDetail(cardId) {
  return {
    type: REQUEST_CARD_DETAIL,
    cardId
  };
}

function receiveCardDetail(cardId, card) {
  return {
    type: RECEIVE_CARD_DETAIL,
    cardId,
    card
  };
}

export const SET_CARD_FILTER_TEXT = 'SET_CARD_FILTER_TEXT';
export function setCardFilterText(key, text) {
  return {
    type: SET_CARD_FILTER_TEXT,
    key,
    text
  };
}

export const SET_CARD_SORT_OPTION = 'SET_CARD_SORT_OPTION';
export function setCardSortOption(key, sort) {
  return {
    type: SET_CARD_SORT_OPTION,
    key,
    sort
  };
}
