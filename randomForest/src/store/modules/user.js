import axios from 'axios'

class User {
  constructor(id) {
    this.id = id
  }
}

const url = 'http://116.203.112.27:8082/v1/user/';

export default {
  state: {
    email: '',
    userId: null,
    isLoggedIn: false,
    loginError: '',
    user: null,
    getUserJson: '',
  },
  mutations: {
    logInUser(state, payload) {
      state.email = payload.email;
      state.userId = payload.userId;
      state.isLoggedIn = true
    },
    loginError(state) {
      state.isLoggedIn = false;
      state.loginError = 'Email и/или Пароль не корректен. Авторизация УПС. '
    },
    setUser(state, payload) {
      state.user = payload
    },
    loadGetUserJson(state, payload) {
      state.getUserJson = payload
    },

  },
  actions: {
    async loginUser({commit}, data) {
      commit('clearError');
      commit('setLoading', true);
      try {
        await axios.post(url + 'login', data,{
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          }
        })
          .then((response) => {
            // console.log(response.data);
            commit('setUser', new User(response.data));
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('loginError');
        throw error
      }
    },
    async registerUser({commit}, data) {
      console.log(data);
      commit('clearError');
      commit('setLoading', true);
      try {
        await axios.post(url + 'register', data)
          .then((response) => {
            console.log(response.data);
            commit('setUser', new User(response.data.id));
            commit('setLoading', false);
          });
      } catch (error) {
        commit('setLoading', false);
        commit('setError', error.message);
        throw error
      }
    },
    async logoutUser({commit}) {
      try {
        await axios.post(url + 'logout')
          .then(commit('setUser', null));

      } catch (error) {
        commit('setError', error.message);
        throw error
      }
    },
    // async autoLoginUser({commit}, id) {
    //   commit('setLoading', true);
    //   try {
    //     await axios.post('http://configurator.talmer.ru/api/session', id)
    //       .then(response => {
    //         console.log(response.data);
    //         commit('setUser', new User(response.data._id));
    //         commit('setLoading', false)
    //       })
    //   } catch (e) {
    //     commit('setError', e.message);
    //     throw e
    //   }
    // },
  },
  getters: {
    isLoggedIn: state => state.isLoggedIn,
    userId: state => state.userId,
    loginError: state => state.loginError,
    isUserLoggedIn(state) {
      return state.user !== null
    },
    user(state) {
      return state.user
    },

    getUser(state) {
      console.log(state.getUserJson)
      return state.getUserJson
    },

  }
}
