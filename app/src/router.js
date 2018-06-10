import Vue from 'vue';
import Router from 'vue-router';
import store from '@/store';
import { AUTH } from '@/store/types';

Vue.use(Router);

// User
// Guest
// Admin

const redirectNonAuthenticatedUsers = (_, __, next) => {
  if (!store.getters[AUTH])
    return next('/');
  return next();
};

const redirectNonGuestUsers = (_, __, next) => {
  if (!!store.getters[AUTH])
    return next('/');
  return next();
};

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/screens/Home')
    },
    {
      path: '/sign-in',
      name: 'SignIn',
      beforeEnter: (_, __, next) => {
        if (user)
          return next('/');
        return next();
      },
      component: () => import('@/screens/SignIn'),
      beforeEnter: redirectNonGuestUsers,
    },
    {
      path: '/sign-on',
      name: 'SignOn',
      component: () => import('@/screens/SignOn'),
      beforeEnter: redirectNonGuestUsers,
    },
    {
      path: '/help',
      name: 'Help',
      component: () => import('@/screens/Help')
    },
  ]
});

export default router;
