<template>
  <nav class="navbar-default navbar-static-side" role="navigation">
    <div class="sidebar-collapse">
      <ul side-navigation class="nav metismenu" id="side-menu">
        <li class="nav-header">
          <div class="profile-element" uib-dropdown>
            <a uib-dropdown-toggle href>
              <span class="clear">
                BigSloth
              </span>
            </a>
          </div>
          <div class="logo-element">
            Sloth
          </div>

        </li>

        <li ui-sref-active="active">
          <a>
            <i class="fa fa-laptop"></i>
            <span class="nav-label">总览</span>
          </a>
        </li>
        <li ui-sref-active="active">
          <a>
            <i class="fa fa-desktop"></i>
            <span class="nav-label">存证</span>
          </a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'main-nav',
  components: {
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  computed: {
    topClass: function () {
    }
  }
}
</script>
