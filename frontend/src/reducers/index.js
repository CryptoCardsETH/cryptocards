import { combineReducers } from 'redux';
import user from './reducer_user';
import card from './reducer_card';
import listing from './reducer_listing';
const rootReducer = combineReducers({
  user,
  card,
  listing
});

export default rootReducer;
