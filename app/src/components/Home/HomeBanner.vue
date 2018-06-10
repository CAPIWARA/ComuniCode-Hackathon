<template>
  <figure class="HomeBanner">
    <impact-button isAlternative size="xl">
      Faça sua<br>Doação
    </impact-button>
  </figure>
</template>

<script>
  import ImpactButton from '@/components/Impact/ImpactButton';

  export default {
    components: { ImpactButton }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .HomeBanner
    display: flex
    align-items: center
    justify-content: center
    width: 100%
    height: calc(100vh - 128px)
    background-color: rgba($color-primary, .8)
</style>
