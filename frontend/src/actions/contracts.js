import apiFetch from './index';
import Eth from 'ethjs-query';
import EthContract from 'ethjs-contract';

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

export function getContractInstance(ABI, contractAddress) {
  const eth = new Eth(window.web3.currentProvider);
  const contract = new EthContract(eth);

  const miniToken = contract(ABI).at(contractAddress);
  return miniToken;
}

export const waitForTxToBeMined = async txHash => {
  const eth = new Eth(window.web3.currentProvider);

  let txReceipt;
  while (!txReceipt) {
    try {
      txReceipt = await eth.getTransactionReceipt(txHash);
    } catch (err) {
      return console.log(err);
    }
  }
  console.log('yayyayayay', txReceipt);
};
