import axios from 'axios'

const url = 'http://116.203.112.27:8082/v1/event/';
// http://116.203.111.141/v1/client/getbycompany/1
export default {
  state: {
    getEvent: '',
    getClientAll: '',
  },
  mutations: {
    loadEvent(state, payload) {
      state.getEvent = payload
    },
    loadClientAll(state, payload) {
      state.getClientAll = payload
    },
  },
  actions: {
    async getEvent({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        await axios.get(url + '/grid/1')
          .then((response) => {
            console.log(response.data);
            commit('loadEvent', response.data);
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    async getClientAll({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        await axios.get('http://116.203.111.141/v1/client/getbycompany/1')
          .then((response) => {
            console.log(response.data);
            commit('loadClientAll', response.data);
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
    getEventJson(state) {
      // console.log(state.getEvent)
      return state.getEvent
    },
    getClientAllJson(state) {
      // console.log(state.getEvent)
      return state.getClientAll
    },
  }
}
