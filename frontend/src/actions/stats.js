import apiFetch from './index';

export const REQUEST_ALL_COUNT_STATS = 'REQUEST_ALL_COUNT_STATS';
export const RECEIVE_ALL_COUNT_STATS = 'RECEIVE_ALL_COUNT_STATS ';

export function fetchAllCountStats() {
  return dispatch => {
    dispatch(requestAllCountStats());
    return apiFetch('stats/counts')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveAllCountStats(json.data));
      });
  };
}

function requestAllCountStats() {
  return {
    type: REQUEST_ALL_COUNT_STATS
  };
}

function receiveAllCountStats(stats) {
  return {
    type: RECEIVE_ALL_COUNT_STATS,
    stats
  };
}
