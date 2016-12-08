<template>
  <div id="login">
    <div class="middle-box text-center loginscreen animated fadeInDown">
      <div>
        <div>
          <h1 class="logo-name">{{logo_name}}</h1>
        </div>
        <div>
          <div class="form-group">
            <input type="email" class="form-control" placeholder="邮箱" required="" v-model="email" >
          </div>
          <div class="form-group">
            <input type="password" class="form-control" placeholder="密码" required="" v-model="password">
          </div>
          <button type="submit" class="btn btn-primary" @click="login">登入</button>
          <router-link :to="{path:'/resetpassword'}"><small>找回密码</small></router-link>

          <button type="submit" class="btn btn-primary" @click="ping">Ping</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {setAccount} from '../../vuex/actions'
import api from '../../api/api'
import Alert from '../../utils/alert'

export default {
  name: 'login',
  data () {
    return {
      logo_name: 'Sloth',
      email: '',
      password: '',
      hi: 'hi'
    }
  },
  computed: {
  },
  methods: {
    login: function () {
      console.log('login.', this.email, this.password)
    },
    ping: function () {
      api.ping().end(function (err, resp) {
        console.log('err', err)
        console.log('resp', resp)
        Alert.success(resp.body.message)
      })
    },
    none: function () {
      setAccount(this.$store)
    }
  }
}
</script>

 <!-- scoped -->
<style lang="scss">
@import '../../scss/style.scss';

body {
  background-color: $gray;
  text-align: center;
}
a {
  cursor: hand;
}
#login {
  text-align: center;
}
.animated {
  margin-right: auto;
  margin-left: auto;
  margin-top: 10%;
  animation-fill-mode: initial;
}
.m-t {
  margin-top: 15px;
}
.form-group {
  margin-bottom: 0;
}
.loginscreen.middle-box {
  width: 300px;
}
.loginColumns {
  max-width: 800px;
  padding: 100px 20px 20px 20px;
}
.passwordBox {
  max-width: 460px;
  padding: 100px 20px 20px 20px;
}
.logo-name {
  color: #e6e6e6;
  font-size: 120px;
  font-weight: 800;
  letter-spacing: -10px;
  margin-bottom: 0;
}
.btn {
  width: 100% !important;
}
.form-control {
  background-color: #FFFFFF;
  background-image: none;
  border: 1px solid #e5e6e7;
  border-radius: 1px;
  color: inherit;
  display: block;
  padding: 6px 12px;
  transition: border-color 0.15s ease-in-out 0s, box-shadow 0.15s ease-in-out 0s;
  width: 100%;
  font-size: 14px
}
</style>