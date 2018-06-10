<template>
  <div class="SignInPill">
    <impact-button
      v-if="!isAuthenticated"
      type="button"
      class="subscribe"
      isAlternative
      @click="$router.push('/sign-on')"
    >Seja um doador</impact-button>

    <impact-button
      v-if="!isAuthenticated"
      type="button"
      size="sm"
      class="login"
      @click="$router.push('/sign-in')"
    >Entrar</impact-button>

    <h4 v-if="isAuthenticated" class="name">{{ user && user.name }}</h4>

    <impact-button
      v-if="isAuthenticated"
      type="button"
      size="sm"
      class="logout"
      @click="logout()"
    >Sair</impact-button>
  </div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex';
  import * as types from '@/store/types';
  import ImpactButton from '@/components/Impact/ImpactButton';

  export default {
    components: { ImpactButton },
    computed: mapGetters({
      user: types.USER,
      isAuthenticated: types.AUTH
    }),
    methods: mapActions({
      logout: types.AUTH_LOGOUT
    })
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .SignInPill
    display: flex
    align-items: center
    justify-content: flex-end

    > .subscribe
      min-width: 90px
      flex-basis: 90px

    > .login
    > .logout
      margin-left: 10px

    > .name
      { $typography-title }

    > .figure
      height: 76px
      margin-right: 12px
</style>
