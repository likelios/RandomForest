import Vue from 'vue'
import App from './App'
import store from './store'
import router from './router'
import Vuetify from 'vuetify'
import * as axios from 'axios'
import 'vuetify/dist/vuetify.min.css'
import LineExample from './components/LineChart.js'
import WindSpeed from './components/WindSpeed.js'
import Humidity from './components/Humidity.js'
import ref from 'vue-ref'
import YmapPlugin from 'vue-yandex-maps'

import Vue2Storage from 'vue2-storage'


const options = { // you may define your apiKey, lang and version or skip this.
  apiKey: '6de7b20a-4375-48d0-8efe-c42fb8a4cbe9', // '' by default
  lang: 'ru_RU', // 'ru_RU' by default
  version: '2.1' // '2.1' by default
};
Vue.use(YmapPlugin, options)
Vue.component('line-example', LineExample);
Vue.component('wind-example', WindSpeed);
Vue.component('humidity-example', Humidity);
Vue.use(Vuetify, {
  theme: {
    primary: '#4d5c7f'
  }
});

Vue.config.productionTip = false;
Vue.use(ref)
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
