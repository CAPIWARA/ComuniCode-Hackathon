<template>
  <div class="HelpOption">
    <img class="image" :alt="title" :src="image" />
    <h3 class="title">{{ title }}</h3>
    <p class="description">{{ description }}</p>
    <impact-button class="button" :to="link" isAlternative>
      {{ button }}
    </impact-button>
  </div>
</template>

<script>
  import ImpactButton from '@/components/Impact/ImpactButton';

  export default {
    components: { ImpactButton },
    props: {
      title: {
        type: String,
        required: true
      },
      description: {
        type: String,
        required: true
      },
      image: {
        type: String,
        required: true
      },
      button: {
        type: String,
        required: true
      },
      link: {
        type: String,
        required: true
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .HelpOption
    display: flex
    flex-direction: column
    align-items: center
    width: 320px
    height: 500px
    padding: 50px
    background-color: rgba($color-primary, .9)

    > .title
      { $typography-title }
      margin-top: 20px
      font-size: 3.5rem
      text-align: center

    > .description
      margin-top: 20px
      text-align: center

    > .image
      filter: invert(100%)

    > .button
      margin-top: auto
</style>

