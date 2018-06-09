<template>
  <nav class="LayoutMenu">
    <router-link
      v-for="(item, index) in items"
      :to="item.path"
      :key="index"
      exact
      active-class="-active"
      exact-active-class="-active"
      class="item"
    >{{ item.name }}</router-link>
  </nav>
</template>

<script>
  import items from '@/items.json';

  export default {
    data: () => ({ items })
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .LayoutMenu
    display: flex
    align-items: center

    > .item
      { $typography-title }
      width: 100px
      height: 50px
      display: inline-flex
      align-items: center
      justify-content: center
      line-height: 1
      color: #000
      text-align: center
      letter-spacing: .25px
      text-decoration: none
      border: 1px dashed transparent
      transition: border-color .8s ease,
                  color .5s ease

    > .item:hover
    > .item.-active
      border-color: $color-primary

    > .item.-active
      color: $color-primary

    > .item + .item
      margin-left: 12px
</style>

