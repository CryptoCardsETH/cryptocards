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

export const REQUEST_TRANSACTION_REPORT = 'REQUEST_TRANSACTION_REPORT';
export const RECEIVE_TRANSACTION_REPORT = 'RECEIVE_TRANSACTION_REPORT';

export function fetchTransactionReport() {
  return dispatch => {
    dispatch(requestTransactionReport());
    return apiFetch('stats/transactionReport')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveTransactionReport(json.data));
      });
  };
}

function requestTransactionReport() {
  return {
    type: REQUEST_TRANSACTION_REPORT
  };
}

function receiveTransactionReport(report) {
  return {
    type: RECEIVE_TRANSACTION_REPORT,
    report
  };
}
