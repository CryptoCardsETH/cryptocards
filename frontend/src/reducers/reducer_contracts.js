import { RECEIVE_CONTRACT_ADDRESSES } from '../actions/contracts';
const INITIAL_STATE = {
  addresses: {},
  loaded: false
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case RECEIVE_CONTRACT_ADDRESSES:
      return {
        ...state,
        addresses: action.contracts,
        loaded: true
      };
    default:
      return state;
  }
}
