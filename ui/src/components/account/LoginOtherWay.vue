<template>
  <div id="login_github">
    Login With Github
    <div>
      <a><i class="fa fa-github-alt fa-4"></i></a>
      <a><i class="fa fa-wechat fa-4"></i></a>
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
  font-size: 80px;
  margin: 10px;
}

#login_github a {
  border: 2px;
  border-color: #444;
}
</style>