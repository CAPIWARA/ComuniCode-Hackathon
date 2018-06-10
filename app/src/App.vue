<template>
  <layout-container>
    <router-view />
  </layout-container>
</template>

<script>
  import { USER } from '@/store/types';
  import LayoutContainer from '@/components/Layout/LayoutContainer';

  export default {
    components: { LayoutContainer },
    mounted () {
      this.$store.dispatch(USER);
    }
  };
</script>

<style lang="stylus">
  @import url('https://fonts.googleapis.com/css?family=Anaheim|Barlow:300,400')
  @import '~@/assets/styles/theme'

  *
    box-sizing: border-box
    margin: 0
    padding: 0

  html
    font-size: 62.5%

  body
    {$typography-base}
</style>
