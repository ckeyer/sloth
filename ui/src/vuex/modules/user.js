import {
  USERINFO_SUCCESS,
  USERINFO_FAILURE,
  SET_ACCOUNT,
  UPDATE_DEVICES
} from '../types'

const state = {
  id: '',
  nicename: '',
  email: '',
  phone: '',
  devices: [],
  passphrase: '',
  state: null,
  logined: false
}

const mutations = {
  [USERINFO_SUCCESS] (state, action) {
    state.state = action.state
  },
  [USERINFO_FAILURE] (state, action) {
    state.state = 'failed state'
  },
  [SET_ACCOUNT] (state, account) {
    if (account !== null) {
      account.logined = true
    }
    // console.log('set account', state)
    for (var k in account) {
      state[k] = account[k]
    }
    state.logined = true
  },
  [UPDATE_DEVICES] (state, device) {
    console.log('update...', device)
    for (var index in state.devices) {
      if (index.id === device.id) {
        state.devices.set(index, device)
        return
      }
    }
    state.devices.push(device)
  }
}

export default {
  state,
  mutations
}