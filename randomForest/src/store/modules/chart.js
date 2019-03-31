import axios from 'axios'

const url = 'http://116.203.112.27:8082/v1/event/1/';
// (Temp; Humidity; WindSpeed)

export default {
  state: {
    tempJson: '',
    windJson: '',
    HumidityJson: ''
  },
  mutations: {
    loadTemp(state, payload) {
      state.tempJson = payload
    },
    loadWind(state, payload) {
      state.windJson = payload
    },
    loadHumidity(state, payload) {
      state.HumidityJson = payload
    },
  },
  actions: {
    async getTemp({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.get(url + 'Temp')
          .then((response) => {
            console.log(response.data);
            commit('loadTemp', response.data);
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    async getWind({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.get(url + 'WindSpeed')
          .then((response) => {
            console.log(response.data);
            commit('loadWind', response.data);
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    async getHumidity({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.get(url + 'Humidity')
          .then((response) => {
            console.log(response.data);
            commit('loadHumidity', response.data);
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
  },
  getters: {
    getTempJson(state) {
      console.log(state.tempJson.Date)
      return state.tempJson
    },
    getWindJson(state) {
      console.log(state.windJson)
      return state.windJson
    },
    getHumidityJson(state) {
      console.log(state.HumidityJson)
      return state.HumidityJson
    }
  }
}
