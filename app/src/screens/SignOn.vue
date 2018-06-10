<template>
  <section class="SignOnScreen">
    <sign-on-form class="form" />
    <figure class="aside">
      <img class="image" src="~@/assets/images/SignOn-Aside.png" />
    </figure>
  </section>
</template>

<script>
  import SignOnForm from '@/components/SignOn/SignOnForm';

  export default {
    components: { SignOnForm }
  };
</script>

<style lang="stylus">
  .SignOnScreen
    display: flex
    flex-direction: column

    > .aside
      height: 100%
      min-height: 460px
      background-size: cover
      background-image: url('~@/assets/images/SignOn-Aside.png')
      background-position: center

      > .image        // background-image da suporte as estratégias de
        display: none // posicionamento e o <img /> nesse caso só serve pra SEO

  @media screen and (min-width: 900px)
    .SignOnScreen
      height: calc(100vh - 98px)
      flex-direction: row

      > .form
      > .aside
        width: 50%
        height: 100%
</style>
