import { RECEIVE_CONTRACT_ADDRESSES } from '../actions/contracts';
const INITIAL_STATE = {
  addresses: {}
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case RECEIVE_CONTRACT_ADDRESSES:
      return {
        ...state,
        addresses: action.contracts
      };
    default:
      return state;
  }
}
