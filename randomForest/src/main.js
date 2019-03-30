import Vue from 'vue'
import App from './App'
import store from './store'
import router from './router'
import Vuetify from 'vuetify'
import * as axios from 'axios'
import 'vuetify/dist/vuetify.min.css'
import LineExample from './components/LineChart.js'

import Vue2Storage from 'vue2-storage'

Vue.component('line-example', LineExample);
Vue.use(Vuetify, {
  theme: {
    primary: '#4d5c7f'
  }
});

Vue.config.productionTip = false;

Vue.use(Vue2Storage, {
  prefix: 'app_',
  driver: 'local',
  ttl: 60 * 60 * 24 * 1000
});


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  axios,
  components: {App},
  template: '<App/>',
  created() {
    console.log(this.$storage);
    if (this.$storage.get('test')) {
      this.$store.dispatch('autoLoginUser', this.$storage.get('test'))
    }

  }
});
