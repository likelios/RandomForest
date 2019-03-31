import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Login from '../components/Auth/Login'
import client from '../components/lk/client'
import admin from '../components/lk/admin'
import Registration from '../components/Auth/Registration'
import AuthGuard from './authGuard'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/admin',
      name: 'admin',
      component: admin
    },
    {
      path: '/registration',
      name: 'reg',
      component: Registration
    },
    {
      path: '/client',
      name: 'client',
      component: client
    },

  ],
  mode: 'history'
})




