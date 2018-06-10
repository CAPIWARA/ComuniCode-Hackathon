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
    height: calc(100vh - 98px)
    display: flex

    > .form
    > .aside
      width: 50%
      height: 100%

    > .aside
      height: 100%
      background-size: cover
      background-image: url('~@/assets/images/SignOn-Aside.png')
      background-position: 0 10%

      > .image        // background-image da suporte as estratégias de
        display: none // posicionamento e o <img /> nesse caso só serve pra SEO
</style>
