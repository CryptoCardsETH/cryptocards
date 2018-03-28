import apiFetch from './index';

export const RECEIVE_CONTRACT_ADDRESSES = 'RECEIVE_CONTRACT_ADDRESSES';
export function fetchContractAddresses() {
  return dispatch => {
    return apiFetch('contracts')
      .then(response => response.json())
      .then(json => {
        dispatch(receiveContractAddresses(json.data));
      });
  };
}
function receiveContractAddresses(contracts) {
  return {
    type: RECEIVE_CONTRACT_ADDRESSES,
    contracts
  };
}
