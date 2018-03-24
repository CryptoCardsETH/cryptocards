import { combineReducers } from 'redux';
import user from './reducer_user';
import card from './reducer_card';
import listing from './reducer_listing';
import stats from './reducer_stats';
const rootReducer = combineReducers({
  user,
  card,
  listing,
  stats
});

export default rootReducer;
