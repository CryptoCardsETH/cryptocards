import {
  REQUEST_ALL_LISTINGS,
  RECEIVE_ALL_LISTINGS
} from '../actions/listings';
import update from 'immutability-helper';
const INITIAL_STATE = {
  all_listings: [],
  all_listings_loading: false
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case REQUEST_ALL_LISTINGS:
      return {
        ...state,
        all_listings_loading: true
      };
    case RECEIVE_ALL_LISTINGS:
      return {
        ...state,
        all_listings: action.listings,
        all_listings_loading: false
      };
    default:
      return state;
  }
}
