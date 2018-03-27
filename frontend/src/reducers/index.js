import { combineReducers } from 'redux';
import user from './reducer_user';
import card from './reducer_card';
import listing from './reducer_listing';
import battle from './reducer_battle';
import stats from './reducer_stats';
const rootReducer = combineReducers({
  user,
  card,
  listing,
  battle,
  stats
});

export default rootReducer;
