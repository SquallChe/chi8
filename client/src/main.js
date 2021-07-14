import Vue from "vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import vuetify from "@/plugins/vuetify";

import Axios from 'axios'

Vue.config.productionTip = false;

//ローカル環境
//Axios.defaults.baseURL = 'http://localhost:8081/';

//本番環境
Axios.defaults.baseURL = 'https://chi8.store:8081/';

Vue.prototype.$http = Axios;

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount("#app");
