<template>
  <section class="LayoutContainer">
    <layout-header class="header" />
    <main class="content">
      <slot />
    </main>
  </section>
</template>

<script>
  import LayoutHeader from '@/components/Layout/LayoutHeader';

  export default {
    components: { LayoutHeader },
  };
</script>

<style lang="stylus">
  .LayoutContainer
    > .header
      position: fixed
      z-index: 1

    > .content
      padding-top: 98px
</style>

