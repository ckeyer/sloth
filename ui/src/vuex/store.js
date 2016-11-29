import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    user,
  },
  strict: false,
  debug: process.env.NODE_ENV !== 'production',
  middlewares
})
