import { REQUEST_ALL_CARDS, RECEIVE_ALL_CARDS } from '../actions/cards';

const INITIAL_STATE = {
  all_cards: [],
  all_cards_loading: false
};

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
    default:
      return state;
  }
}
