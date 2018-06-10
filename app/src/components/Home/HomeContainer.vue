<template>
  <section class="HomeContainer">
    <figure class="aside">
      <img class="image" src="~@/assets/images/Aside.png" />
    </figure>
    <home-content class="content" />
  </section>
</template>

<script>
  import HomeContent from '@/components/Home/HomeContent'

  export default {
    components: { HomeContent }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .HomeContainer
    display: flex
    flex-direction: column

    > .aside
      background-size: cover
      background-image: url('~@/assets/images/Aside.png')
      background-position: 0 10%
      min-height: 460px

      > .image        // background-image da suporte as estratégias de
        display: none // posicionamento e o <img /> nesse caso só serve pra SEO

  @media screen and (min-width: 768px)
    .HomeContainer
      flex-direction: row
      height: 900px

      > .aside
        width: 50%

      > .content
        width: 50%

  @media screen and (min-width: 1200px)
    .HomeContainer
      height: 900px

      > .aside
        width: 30%

      > .content
        width: 70%
</style>
