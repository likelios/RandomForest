import axios from 'axios'

const url = 'http://116.203.112.27:8082/v1/event/';
export default {
  state: {
    getEvent: '',
    getClientAll: '',
    getNonClientAll: '',
  },
  mutations: {
    loadEvent(state, payload) {
      state.getEvent = payload
    },
    loadClientAll(state, payload) {
      state.getClientAll = payload
    },
    loadNonClientAll(state, payload) {
      state.getNonClientAll = payload
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
          });
        await axios.get('http://116.203.112.27:8082/v1/client/getbynotcompany/1')
          .then((response) => {
            console.log(response.data);
            commit('loadNonClientAll', response.data);
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    // async getnonClientAll({commit}) {
    //   commit('clearError');
    //   commit('setLoading', true);
    //   try {
    //     await axios.get('http://116.203.112.27:8082/v1/client/getbynotcompany/1')
    //       .then((response) => {
    //         console.log(response.data);
    //         commit('loadNonClientAll', response.data);
    //         commit('setLoading', false);
    //       });
    //   } catch (error) {
    //     commit('setLoading', false);
    //     commit('loginError');
    //     throw error
    //   }
    // },
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
    getNonClientAllJson(state) {
      // console.log(state.getEvent)
      return state.getNonClientAll
    },
  }
}
