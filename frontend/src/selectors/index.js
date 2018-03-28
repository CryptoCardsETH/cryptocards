import { createSelector } from 'reselect';
import { ABIs } from '../contracts';

import Eth from 'ethjs-query';
import EthContract from 'ethjs-contract';

const getUserState = state => state.user;
const getContractsState = state => state.contract;

export const isReadyForContract = createSelector(
  [getUserState, getContractsState],
  (user, contract) => {
    console.log(contract);
    return user.me.id && contract.loaded;
  }
);

export const getContractAddress = (state, contractName) => {
  let ready = isReadyForContract(state);
  let contract = getContractsState(state);
  return ready ? contract.addresses[contractName].address : null;
};

export const getContractByName = (state, contractName) => {
  return {
    abi: ABIs[contractName],
    address: getContractAddress(state, contractName)
  };
};

export const getContractInstanceByName = (state, contractName) => {
  const info = getContractByName(state, contractName);
  return getContractInstance(info.abi, info.address);
};

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
