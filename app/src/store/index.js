import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import authentication from '@/store/modules/authentication';
import user from '@/store/modules/user';

Vue.use(Vuex);

const store = new Store({
  strict: true,
  modules: {
    authentication,
    user
  }
});

export default store;
