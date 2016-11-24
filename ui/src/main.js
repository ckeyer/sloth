import Vue from 'vue'

import router from './router'
import App from './App'

// CSS
import 'font-awesome/css/font-awesome.css'
import 'bootstrap/dist/css/bootstrap.css'
import 'vue-strap/dist/vue-strap'

/* eslint-disable no-new */
new Vue({
  router,
  el: '#app',
  template: '<App/>',
  components: { App }
})
