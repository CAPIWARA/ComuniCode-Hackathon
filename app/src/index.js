import Vue from 'vue';
import App from '@/App';
import store from '@/store';
import router from '@/router';
import Validate, { Validator } from 'vee-validate';
import ValidatePortuguese from 'vee-validate/dist/locale/pt_BR';

Validator.localize('pt_BR', ValidatePortuguese);

Vue.use(Validate, {
  locale: 'pt_BR',
  dictionary: {
    sv: ValidatePortuguese
  }
});

new Vue({
  el: '#app',
  store,
  router,
  render: (λ) => λ(App)
});
