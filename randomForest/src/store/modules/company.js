import axios from 'axios'

const url = 'http://116.203.111.141/api/v1/';


export default {
  state: {
    companyJson: ''
  },
  mutations: {
    loadCompany(state, payload) {
      state.companyJson = payload
    },
  },
  actions: {
    async getComponyJson({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.post(url + 'companies')
          .then((response) => {
            console.log(response.data);
            commit('loadCompany', response.data);
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
    companyJson(state) {
      return state.companyJson.companies
    }
  }
}
