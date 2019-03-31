import Vue from 'vue'
import Vuex from 'vuex'
import shared from './shared'
import user from './modules/user'
import company from './modules/company'
import chart from './modules/chart'
import data from './modules/data'


Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    shared,
    user,
    company,
    chart,
    data,
  }
})
