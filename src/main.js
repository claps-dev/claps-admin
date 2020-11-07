import Vue from 'vue';
import Vuex from 'vuex';
import VueAxios from "vue-axios";
import axios from "axios";
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;

// axios
Vue.use(VueAxios, axios);

// vuex
Vue.use(Vuex);

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app');
