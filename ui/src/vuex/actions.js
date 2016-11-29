import {
  // USERINFO_SUCCESS,
  // USERINFO_FAILURE,
  SET_ACCOUNT
  // UPDATE_DEVICES
} from './types'

export const setAccount = ({ dispatch }, account) => {
  dispatch(SET_ACCOUNT, account)
}
