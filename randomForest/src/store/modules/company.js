import axios from 'axios'

const url = 'http://116.203.112.27:8082/v1/';

// http://localhost:8080/v1/company/{ID
export default {
  state: {
    companyJson: '',
    companyOneJson: ''
  },
  mutations: {
    loadCompany(state, payload) {
      state.companyJson = payload
    },
    loadCompanyOne(state, payload) {
      state.companyOneJson = payload
    },
  },
  actions: {
    async getComponyJson({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.get(url + 'company')
          .then((response) => {
            // console.log(response.data);
            commit('loadCompany', response.data);
            commit('setLoading', false);

          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    async getComponyOneJson({commit}) {
      commit('clearError');
      commit('setLoading', true);
      try {
        axios.get(url + 'company/1')
          .then((response) => {
            // console.log(response.data);
            commit('loadCompanyOne', response.data);
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
      // console.log(state.companyJson)
      return state.companyJson
    },
    companyOnesJson(state) {
      // console.log(state.companyJson)
      return state.companyOneJson
    }
  }
}
