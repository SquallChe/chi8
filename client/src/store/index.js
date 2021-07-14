import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    items:[]
  },
  getters: {
    getItems: state => {
      return state.items
    }
  },
  mutations: {
    setItems(state, items) {
      state.items = items
    }
  },
  actions: {},
  modules: {}
});
