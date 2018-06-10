<template>
  <section class="SignInScreen">
    <sign-in-form class="form" @submit="login" />
    <figure class="aside">
      <img class="image" src="~@/assets/images/SignOn-Aside.png" />
    </figure>
    <p v-if="error" class="error">{{ error }}</p>
  </section>
</template>

<script>
  import SignInForm from '@/components/SignIn/SignInForm';
  import { AUTH_LOGIN } from '@/store/types';
  import { mapActions } from 'vuex';

  export default {
    components: { SignInForm },
    data: () => ({
      error: ''
    }),
    methods: {
      async login (params) {
        try {
          await this.$store.dispatch(AUTH_LOGIN, params);
          this.error = '';
          this.$router.push('/help');
        } catch (error) {
          console.dir(error);
          this.error = 'Erro ao se autenticar. E-Mail e senha de usuário podem não coincidir';
        }
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .SignInScreen
    display: flex
    flex-direction: column

    > .error
      color: $color-error

    > .aside
      height: 100%
      min-height: 460px
      background-size: cover
      background-image: url('~@/assets/images/SignOn-Aside.png')
      background-position: center

      > .image        // background-image da suporte as estratégias de
        display: none // posicionamento e o <img /> nesse caso só serve pra SEO

  @media screen and (min-width: 900px)
    .SignInScreen
      height: calc(100vh - 98px)
      flex-direction: row

      > .form
      > .aside
        width: 50%
        height: 100%
</style>
