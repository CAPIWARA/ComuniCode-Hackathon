<template>
  <component
    v-bind="$attrs"
    v-on="$listeners"
    :class="classes"
    :is="to ? 'router-link' : 'button'"
    :to="to ? to : null"
  >
    <span class="text">
      <slot />
    </span>
  </component>
</template>

<script>
  export default {
    props: {
      to: {
        type: [Object, String],
        default: null
      }
    },
    computed: {
      classes () {
        return [
          'ImpactButton', {
            '-is-link': this.to,
            '-is-button': !this.to
          }
        ]
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .ImpactButton
    display: flex
    align-items: center
    justify-content: center
    height: 48px
    padding-left: 12px
    padding-right @padding-left
    border: 2px solid $color-primary
    outline: none
    cursor: pointer
    transition: background-color .2s ease-in

    > .text
      {$typography-title}
      font-size: 1.6rem
      text-align: center
      text-transform: uppercase
      letter-spacing: 1px
      color: $color-primary
      transition: color .1s ease-in

    &.-is-button
      background-image: none
      background-color: transparent

    &.-is-link
      text-decoration: none

    &:hover
      background-color: $color-primary

      > .text
        color: #FFFFFF
</style>
