<template>
  <form class="ImpactForm" novalidate @submit.prevent="$emit('submit')">
    <slot />
  </form>
</template>

<style lang="stylus">
  .ImpactForm
    display: flex
    flex-direction: column
    align-items: flex-start
    justify-content: center
    padding: 48px

    .ImpactEntry + .ImpactEntry
      margin-top: 24px
</style>
