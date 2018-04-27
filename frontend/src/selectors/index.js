import { createSelector } from 'reselect';
import { ABIs, CONTRACT_NAME_CORE } from '../contracts';

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

export const getContractAddress = async (state, contractName) => {
  let ready = isReadyForContract(state);
  if (!ready) return null;

  let contract = getContractsState(state);
  if (contractName === CONTRACT_NAME_CORE)
    return contract.addresses[contractName].address;
  else {
    console.log('fuck', contractName);
    //grab the address from core
    let core = await getContractInstanceByName(state, CONTRACT_NAME_CORE);
    let fullName;
    if (contractName === 'BattleGroups') fullName = 'BattleGroupContract';
    if (contractName === 'BattleQueue') fullName = 'BattleQueueContract';
    let res = await core[fullName].call();
    return res[0];
  }
};

export const getContractByName = async (state, contractName) => {
  let address = await getContractAddress(state, contractName);
  return {
    abi: ABIs[contractName],
    address
  };
};

export const getContractInstanceByName = async (state, contractName) => {
  const info = await getContractByName(state, contractName);
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
  console.log('Transaction Sent!', txHash);
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
