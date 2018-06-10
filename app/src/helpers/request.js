import store from '@/store';
import { AUTH_TOKEN } from '@/store/types';

const API = 'http://172.16.3.46:8080';

const request = async (method, url, params = {}) => {
  const response = await fetch(API + url, {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params)
  });

  return response;
};

const requestGraphQL = async (query = '', variables = {}) => {
  const token = store.getters[AUTH_TOKEN] || null;
  const response = await fetch(API + '/graphql', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ... token && {
        'Authorization': token
      }
    },
    body: JSON.stringify({
      query,
      variables
    })
  });

  return response.json();
};

export { URL, request, requestGraphQL };
