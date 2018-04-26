import apiFetch from './index';
import { toast } from 'react-toastify';

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

export function sellCards(cardIds) {
  console.log('cardss');
  console.log(cardIds);
  return dispatch => {
    return apiFetch('sell', {
      method: 'PUT',
      body: JSON.stringify(cardIds)
    })
      .then(response => response.json())
      .then(json => {
        if (json.success) {
          toast.success('Cards successfully added to listings!');
        } else toast.error('Failure adding cards to listing');
      });
  };
}
