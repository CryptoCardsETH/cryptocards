import apiFetch, { persistLocally } from './index';

export const RECEIVE_ALL_BATTLES = 'RECEIVE_ALL_BATTLES';

export function fetchAllBattles() {
  return dispatch => {
    return apiFetch('battles')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveAllBattles(json.data));
      });
  };
}

function receiveAllBattles(battles) {
  return {
    type: RECEIVE_ALL_BATTLES,
    battles
  };
}
