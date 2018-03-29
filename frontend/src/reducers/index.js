import { combineReducers } from 'redux';
import user from './reducer_user';
import card from './reducer_card';
import listing from './reducer_listing';
import battle from './reducer_battle';
import stats from './reducer_stats';
import contract from './reducer_contracts';
const rootReducer = combineReducers({
  user,
  card,
  listing,
  contract,
  battle,
  stats
});

export default rootReducer;
