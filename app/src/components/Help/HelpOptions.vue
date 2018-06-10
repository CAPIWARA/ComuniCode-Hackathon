<template>
  <section class="HelpOptions">
    <help-option
      v-for="(option, index) in options"
      :key="index"
      :link="option.link"
      :title="option.title"
      :image="option.image"
      :button="option.button"
      :description="option.description"
      class="option"
    />
  </section>
</template>

<script>
  import options from '@/options';
  import HelpOption from '@/components/Help/HelpOption';

  export default {
    components: { HelpOption },
    data: () => ({ options })
  }
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .HelpOptions
    width: 100%
    height: 100%
    display: flex
    align-items: center
    justify-content: center

    > .option:nth-child(1)
      margin-right: -30px

    > .option:nth-child(3)
      margin-left: -30px

    > .option:nth-child(2)
      width: 360px
      height: 560px
      padding: 60px
      background-color: $color-primary
</style>

