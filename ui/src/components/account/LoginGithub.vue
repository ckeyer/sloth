<template>
  <div id="login_github">
    Login With Github
    <div>
      <a><i class="fa fa-github-alt fa-6"></i></a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'login-github',
  ready: function () {
    console.log('hello github.')
  }
}
</script>

<style scope>
#login_github i {
  font-size: 240px;
}
</style>