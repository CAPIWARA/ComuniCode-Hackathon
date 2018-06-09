import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

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
      component: () => import('@/screens/SignIn')
    }
  ]
});

export default router;
