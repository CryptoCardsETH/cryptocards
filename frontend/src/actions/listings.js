import apiFetch from './index';

export const REQUEST_ALL_LISTINGS = 'REQUEST_ALL_LISTINGS';
export const RECEIVE_ALL_LISTINGS = 'RECEIVE_ALL_LISTINGS';

export function fetchAllListings() {
  return dispatch => {
    dispatch(requestAllListings());
    return apiFetch('listings')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveAllListings(json.data));
      });
  };
}

function requestAllListings() {
  return {
    type: REQUEST_ALL_LISTINGS
  };
}

function receiveAllListings(listings) {
  return {
    type: RECEIVE_ALL_LISTINGS,
    listings
  };
}
