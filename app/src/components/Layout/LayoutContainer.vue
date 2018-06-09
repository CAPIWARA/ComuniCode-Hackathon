<template>
  <section>
    <layout-header />
    <slot />
  </section>
</template>

<script>
  import LayoutHeader from '@/components/Layout/LayoutHeader';

  export default {
    components: { LayoutHeader },
  };
</script>
