import * as types from '@/store/types';
import decode from 'jwt-decode';

const storage = window.localStorage;

const state = {
  token: storage.getItem('token')
};

const getters = {
  [types.AUTH]: (state) => {
    try {
      const { exp: expiration } = decode(state.token);
      const current = Date.now() / 1000;
      const isValid = !expiration || +expiration > current;
      return isValid;
    } catch (_) {
      return false;
    }
  }
};

const mutations = {
  [types.AUTH_TOKEN]: (state, token) => {
    state.token = token;
    if (!token)
      storage.removeItem('token');
    else
      storage.setItem('token', token);
  }
};

const actions = {
  [types.AUTH_LOGIN]: async ({ state, commit }, params) => {
    const response = await fetch('https://', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: params.email,
        password: params.password
      }),
    });
    const token = response.headers.get('x-amzn-remapped-authorization');
    commit(types.AUTH, token);
  },

  [types.AUTH_LOGOUT]: ({ commit }) => {
    commit(types.AUTH_TOKEN, null);
  }
};

export default { state, getters, mutations, actions };
