import {
  REQUEST_ALL_CARDS,
  RECEIVE_ALL_CARDS,
  RECEIVE_CARD_DETAIL,
  SET_CARD_FILTER_TEXT,
  SET_CARD_SORT_OPTION
  EDIT_CARD_DETAIL
} from '../actions/cards';
import update from 'immutability-helper';
const INITIAL_STATE = {
  all_cards: [],
  all_cards_loading: false,
  card_detail: {},
  card_detail_loading: false,
  filters: {}
};

update.extend('$auto', function(value, object) {
  return object ? update(object, value) : update({}, value);
});

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case REQUEST_ALL_CARDS:
      return {
        ...state,
        all_cards_loading: true
      };
    case RECEIVE_ALL_CARDS:
      return {
        ...state,
        all_cards: action.cards,
        all_cards_loading: false
      };
    case RECEIVE_CARD_DETAIL:
      return update(state, {
        card_detail: {
          [action.cardId]: { $set: action.card }
        }
      });
    case SET_CARD_FILTER_TEXT:
      return update(state, {
        filters: {
          [action.key]: { $auto: { text: { $set: action.text } } }
        }
      });
    case SET_CARD_SORT_OPTION:
      return update(state, {
        filters: {
          [action.key]: { $auto: { sort: { $set: action.sort } } }
    case EDIT_CARD_DETAIL:
      return update(state, {
        card_detail: {
          [action.cardId]: {
            [action.key]: { $set: action.value }
          }
        }
      });
    default:
      return state;
  }
}
