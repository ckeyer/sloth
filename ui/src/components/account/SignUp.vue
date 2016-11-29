<template>
  <div id="signup">
    SignUp
    <div class="form-group">
      <label class="col-sm-2 control-label">api root path</label>
      <div class="col-sm-10">{{apiPath}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'signup',
  data () {
    return {
      apiPath: process.env.API_ROOT
    }
  }
}
</script>

<style lang="scss">

</style>
