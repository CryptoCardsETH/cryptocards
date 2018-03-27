import {
  REQUEST_ALL_COUNT_STATS,
  RECEIVE_ALL_COUNT_STATS,
  REQUEST_TRANSACTION_REPORT,
  RECEIVE_TRANSACTION_REPORT
} from '../actions/stats';

const INITIAL_STATE = {
  count_stats: {},
  count_stats_loading: false,

  transaction_report: {},
  transaction_report_loading: false
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case REQUEST_ALL_COUNT_STATS:
      return {
        ...state,
        count_stats_loading: true
      };
    case RECEIVE_ALL_COUNT_STATS:
      return {
        ...state,
        count_stats: action.stats,
        count_stats_loading: false
      };
    case REQUEST_TRANSACTION_REPORT:
      return {
        ...state,
        transaction_report_loading: true
      };
    case RECEIVE_TRANSACTION_REPORT:
      return {
        ...state,
        transaction_report: action.report,
        transaction_report_loading: false
      };
    default:
      return state;
  }
}
