<template>
  <header class="LayoutHeader">
    <div class="container">
      <img src="~@/assets/images/Logo.png" class="figure" alt="" />

      <layout-menu />

      <impact-button
        type="button"
        class="subscribe"
        isAlternative
        @click="$router.push('/sign-on')"
      >
        Seja um doador
      </impact-button>

      <impact-button
        type="button"
        size="sm"
        class="login"
        @click="$router.push('/sign-in')"
      >
        Entrar
      </impact-button>
    </div>
  </header>
</template>

<script>
  import ImpactButton from '@/components/Impact/ImpactButton';
  import LayoutMenu from '@/components/Layout/LayoutMenu';

  export default {
    components: { LayoutMenu, ImpactButton }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .LayoutHeader
    width: 100%
    height: 98px
    padding: 0 12px
    border-bottom: 2px solid $color-primary
    background-color: #FFFFFF

    > .container
      width: 100%
      height: 100%
      max-width: 1050px
      display: flex
      align-items: center
      justify-content: flex-start
      margin-left: auto
      margin-right: auto

      > .subscribe
        margin-left: auto

      > .login
        margin-left: 10px

      > .figure
        height: 76px
        margin-right: 12px
</style>

