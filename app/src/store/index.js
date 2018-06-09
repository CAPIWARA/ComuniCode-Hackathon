import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import authentication from '@/store/modules/authentication';

Vue.use(Vuex);

const store = new Store({
  strict: true,
  modules: {
    authentication
  }
});

export default store;
