import { combineReducers } from 'redux';
import user from './reducer_user';
import card from './reducer_card';
import listing from './reducer_listing';
import transaction from './reducer_card';
const rootReducer = combineReducers({
  user,
  card,
  listing,
  transaction
});

export default rootReducer;
