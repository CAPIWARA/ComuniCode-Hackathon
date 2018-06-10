<template>
  <fieldset :class="classes">
    <label class="label">{{ label }}</label>
    <input
      v-bind="$attrs"
      :name="name"
      :value="value"
      class="entry"
      @input="$emit('input', $event.target.value)"
    />
    <span class="error" v-html="error || '&nbsp;'"></span>
  </fieldset>
</template>

<script>
  export default {
    props: {
      name: {
        type: String,
        required: true
      },
      value: {
        type: String,
        required: true
      },
      error: {
        type: String,
        default: null
      },
      label: {
        type: String,
        default: null
      }
    },
    computed: {
      classes () {
        return [
          'ImpactEntry', {
            '-is-error': !!this.error
          }
        ];
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .ImpactEntry
    display: flex
    flex-direction: column
    width: 100%
    border: 0 none transparent

    > .label
      font-size: 1.6rem
      margin-left: 10px
      margin-bottom: 6px

    > .entry
      { $typography-title }
      width: 100%
      height: 36px
      border: 2px solid #000000
      padding-left: 10px
      padding-right @padding-left
      font-size: 1.4rem
      outline: none
      transition: border-bottom-color .3s ease
      caret-color: $color-primary

    > .entry:focus
      border-bottom-color: $color-primary

    > .error
      color: $color-error

    &.-is-error > .entry
      border-bottom-color: $color-error
      animation: shake .8s ease

  $shake-angle = 0

  @keyframes shake
    0
      transform: rotate($shake-angle + 1deg)
    10%
      transform: rotate($shake-angle + -1deg)
    10%
      transform: rotate($shake-angle + 1deg)
    30%
      transform: rotate($shake-angle + -1deg)
    40%
      transform: rotate($shake-angle + 1deg)
    50%
      transform: rotate($shake-angle + -.5deg)
    60%
      transform: rotate($shake-angle + .25deg)
    60%
      transform: rotate(-.125deg)
</style>

