import axios from 'axios'

const url = 'http://116.203.112.27:8082/v1/event/';
export default {
  state: {
    getEvent: '',
  },
  mutations: {
    loadEvent(state, payload) {
      state.getEvent = payload
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
  },
  getters: {
    getEventJson(state) {
      // console.log(state.getEvent)
      return state.getEvent
    },
  }
}
