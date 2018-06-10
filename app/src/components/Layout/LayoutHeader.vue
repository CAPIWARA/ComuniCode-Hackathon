<template>
  <header class="LayoutHeader">
    <div class="container">
      <img src="~@/assets/images/Logo-Wide.png" class="figure" alt="" />

      <layout-menu />

      <sign-in-pill class="sign-in" />
    </div>
  </header>
</template>

<script>
  import ImpactButton from '@/components/Impact/ImpactButton';
  import LayoutMenu from '@/components/Layout/LayoutMenu';
  import SignInPill from '@/components/SignIn/SignInPill';

  export default {
    components: { LayoutMenu, ImpactButton, SignInPill },
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

      > .sign-in
        margin-left: auto

      > .figure
        height: 76px
        margin-right: 12px
</style>

