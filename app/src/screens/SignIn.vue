<template>
  <section class="SignInScreen">
    <sign-in-form class="form" @submit="login" />
    <figure class="aside">
      <img class="image" src="~@/assets/images/SignOn-Aside.png" />
    </figure>
  </section>
</template>

<script>
  import SignInForm from '@/components/SignIn/SignInForm';
  import { AUTH_LOGIN } from '@/store/types';
  import { mapActions } from 'vuex';

  export default {
    components: { SignInForm },
    methods: mapActions({
      login: AUTH_LOGIN
    })
  };
</script>

<style lang="stylus">
  .SignInScreen
    display: flex
    flex-direction: column

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
