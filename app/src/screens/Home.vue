<template>
  <section class="HomeScreen">
    <figure class="background"></figure>
    <home-banner />
    <home-container />
  </section>
</template>

<script>
  import HomeBanner from '@/components/Home/HomeBanner';
  import HomeContainer from '@/components/Home/HomeContainer';

  export default {
    components: { HomeBanner, HomeContainer }
  };
</script>

<style lang="stylus">
  .HomeScreen
    position: relative

    > .background
      position: fixed
      z-index: -1
      width: 100%
      height: 100vh
      background-image: url('~@/assets/images/Home-Background.jpg')
      background-position: center
      background-size: cover
</style>


