<template>
  <section class="HelpScreen">
    <help-options />
  </section>
</template>

<script>
  import HelpOptions from '@/components/Help/HelpOptions';

  export default {
    components: { HelpOptions },
  }
</script>

<style lang="stylus">
  .HelpScreen
    background-image: url('~@/assets/images/Help-Background.png')
    height: calc(100vh - 98px)
</style>
