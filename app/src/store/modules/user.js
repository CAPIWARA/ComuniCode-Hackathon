import * as types from '@/store/types';
import { requestGraphQL } from '@/helpers/request';

const state = {
  user: null
};

const getters = {
  [types.USER]: (state) => state.user
};

const mutations = {
  [types.USER]: (state, user) => {
    state.user = user;
  }
};

const actions = {
  [types.USER]: async ({ commit, getters }) => {
    const token = getters[types.AUTH_TOKEN] || null;

    if (!token)
      return

    const { data: user } = await requestGraphQL(`
      query {
        user {
          name
          email
        }
      }
    `);

    commit(types.USER, user);
  }
};

export default { state, getters, mutations, actions };
