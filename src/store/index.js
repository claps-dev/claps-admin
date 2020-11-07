import Vue from 'vue'
import Vuex from 'vuex'
import userModule from './module/user'

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
  },
  mutations: {
  },

  // 开启严格模式
  // strict: process.env.NODE_ENV !== "production",

  actions: {
  },
  modules: {
    userModule: userModule,
  },
});

export default store
