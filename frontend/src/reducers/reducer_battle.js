import { RECEIVE_ALL_BATTLES } from '../actions/battles';
const INITIAL_STATE = {
  all_battles: []
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case RECEIVE_ALL_BATTLES:
      return {
        ...state,
        all_battles: action.battles
      };
    default:
      return state;
  }
}
