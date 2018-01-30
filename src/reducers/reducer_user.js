import {
  LOGIN_FROM_JWT_SUCCESS,
  RECEIVE_ME,
  SET_SIGNED_MESSAGE,
  SET_WEB3_AVAILABILITY
} from '../actions/users';
const INITIAL_STATE = {
  authenticated: false,
  me: {
    email: ''
  },
  isWeb3Available: false,
  signedMessage: '',
  jwt: null
};

export default function(state = INITIAL_STATE, action) {
  switch (action.type) {
    case SET_WEB3_AVAILABILITY:
      return {
        ...state,
        isWeb3Available: action.isAvailable
      };
    case SET_SIGNED_MESSAGE:
      return {
        ...state,
        signedMessage: action.message
      };
    case LOGIN_FROM_JWT_SUCCESS:
      return {
        ...state,
        authenticated: true,
        jwt: action.token
      };
    case RECEIVE_ME:
      //todo: error checking
      return {
        ...state,
        me: action.me
      };
    default:
      return state;
  }
}
