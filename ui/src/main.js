import Vue from 'vue'

import router from './router'
import App from './App'

// CSS
import 'font-awesome/css/font-awesome.css'
import 'bootstrap/dist/css/bootstrap.css'
import 'vue-strap/dist/vue-strap'
import 'toastr/package/build/toastr.min.css'

const debug = process.env.NODE_ENV !== 'production'

Vue.config.debug = debug
Vue.config.warnExpressionErrors = false

/* eslint-disable no-new */
new Vue({
  router,
  el: '#app',
  template: '<App/>',
  components: { App }
})
