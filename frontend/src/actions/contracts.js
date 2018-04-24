import apiFetch from './index';

export const RECEIVE_CONTRACT_ADDRESSES = 'RECEIVE_CONTRACT_ADDRESSES';
export const fetchContractAddresses = () => dispatch =>
  apiFetch('contracts')
    .then(response => response.json())
    .then(json => {
      dispatch(receiveContractAddresses(json.data));
    });

const receiveContractAddresses = contracts => ({
  type: RECEIVE_CONTRACT_ADDRESSES,
  contracts
});
