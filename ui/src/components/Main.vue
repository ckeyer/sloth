<template>
  <div  id="wrapper" class="main pace-done" :class="[topClass]">
    <main-nav></main-nav>
    <div id="page-wrapper" class="gray-bg">
        <!-- Page wrapper -->
        <top-nav-bar></top-nav-bar>

        <!-- Main view  -->
        <router-view></router-view>

        <!-- Footer -->
        <main-footer></main-footer>
    </div>
  </div>
</template>

<script>
import MainNav from './common/MainNav'
import TopNavBar from './common/TopNavBar'
import MainFooter from './common/MainFooter'

import 'jquery/dist/jquery.slim.min'
import 'jquery/dist/jquery.min'
import 'pace/pace'
// import '../utils/inspinia'

export default {
  name: 'main',
  components: {
    MainNav,
    TopNavBar,
    MainFooter
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  computed: {
    topClass: function () {
      console.log('this.', this.$el)
      // if (this.$el.clientWidth < 700) {
      //   return 'body-small'
      // }
      return ''
    }
  }
}
</script>

