<template>
  Login Whit Username.

  <div>
    <div class="form-group">
      <input type="email" class="form-control" placeholder="邮箱" required="" v-model="email" >
    </div>
    <div class="form-group">
      <input type="password" class="form-control" placeholder="密码" required="" v-model="password">
    </div>
    <button type="submit" class="btn btn-primary" :class="{disabled: !allowSubmit}" @click="login">登入</button>
    <button type="submit" class="btn btn-primary" @click="ping">Ping</button>
    <router-link :to="{path:'/resetpassword'}"><small>找回密码</small></router-link>
  </div>
</template>

<script>
import {setAccount} from '../../vuex/actions'
import api from '../../api/api'
import Alert from '../../utils/alert'

export default {
  name: 'login-user',
  data () {
    return {
      logo_name: 'Sloth',
      email: '',
      password: ''
    }
  },
  created: function () {
    console.log('hello', this)
    console.log('router', this.$route)
  },
  computed: {
    allowSubmit: function () {
      return this.email !== '' && this.password.length > 4
    }
  },
  methods: {
    login: function () {
      let body = {
        email: this.email,
        password: this.password
      }
      console.log('login.', body)
      api.login(body).end(function (err, resp) {
        if (err) {
          Alert.error(err, '请求错误')
        }
        console.log('post login', err, resp)
      })
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
