import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from './components/account/Login'
import LoginOtherWay from './components/account/LoginOtherWay'
import LoginUser from './components/account/LoginUser'
import SignUp from './components/account/SignUp'
import ResetPassword from './components/account/ResetPassword'

import Main from './components/Main'

Vue.use(VueRouter)

const Foo = Vue.extend({
  template: '<p>This is foo!</p>'
})

const router = new VueRouter({
  mode: 'history',
  // base: __dirname,
  routes: [
    {
      path: '/login',
      component: Login,
      children: [
        {
          path: 'others',
          component: LoginOtherWay
        },
        {
          path: '',
          component: LoginUser
        }
      ]
    },
    {
      path: '/signup',
      component: SignUp
    },
    {
      path: '/resetpassword',
      component: ResetPassword
    },
    {
      path: '/',
      component: Main,
      children: [
        {
          path: 'foo',
          component: Foo
        }
      ]
    }
  ]
})

export default router
