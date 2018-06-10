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
      },
      isAlternative: {
        type: Boolean,
        default: false
      },
      size: {
        type: String,
        default: 'md',
        validator: (size) => [ 'xl', 'md', 'sm' ].includes(size)
      }
    },
    computed: {
      classes () {
        return [
          'ImpactButton',
          '-' + this.size, {
            '-is-link': this.to,
            '-is-button': !this.to,
            '-is-alternative': this.isAlternative
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
    border: 2px solid $color-primary
    outline: none
    cursor: pointer
    transition: background-color .2s ease-in

    > .text
      {$typography-title}
      text-align: center
      text-transform: uppercase
      letter-spacing: 1px
      color: $color-primary
      line-height: 1
      transition: color .1s ease-in

    &.-is-button
      background-image: none
      background-color: transparent

    &.-is-link
      text-decoration: none

    &.-is-alternative
      border-color: #000000

      > .text
        color: #000000

      &:hover
        background-color: #000000

    &:hover
      background-color: $color-primary

      > .text
        color: #FFFFFF

    &.-xl
      height: 150px
      padding-left: 50px
      padding-right @padding-left

      > .text
        font-size: 3rem

    &.-md
      height: 48px
      padding-left: 12px
      padding-right @padding-left

      > .text
        font-size: 1.6rem

    &.-sm
      height: 36px
      padding-left: 10px
      padding-right @padding-left

      > .text
        font-size: 1.4rem
</style>
