<template>
  <section class="HomeScreen">
    <figure class="aside">
      <img class="image" src="~@/assets/images/Aside.png" />
    </figure>
    <section class="HomeContent">
      <article class="content">
        <h2 class="title">Quem somos</h2>
        <p class="paragraph">O Círculo de Amigos da Paróquia de Santa Madre Cabrini (razão social), fundado em 1984, é uma sociedade beneficente de direito privado, sem fins lucrativos, reconhecida de utilidade pública municipal, estadual e federal, cadastrada no Conselho Nacional de Assistência Social (CNAS).</p>

        <p class="paragraph">O “Projeto Madre Cabrini” (nome fantasia), atualmente atende 120 menores carentes, 06 a 14 anos, em suas necessidades mais urgentes, realizando atividades de complementação a escola regular, como:  apoio pedagógico, informática, artes  e artesanato, esportes e recreação, refeições, formação moral e religiosa, música, dança, culinária, horta e jardinagem, acompanhamento psicológico, atendimento odontológico, pintura em tela, dentre outros.</p>
      </article>

      <article class="content">
        <h2 class="title">O projeto</h2>
        <p class="paragraph">O Círculo de Amigos da Paróquia de Santa Madre Cabrini foi idealizado em 1984 com o início da urbanização da antiga favela do Gonzaga, quando a comunidade assumiu a ação de acolher os menores carentes. E, enquanto os pais trabalhavam em sistema de mutirão, as crianças e adolescentes recebiam atendimento no salão paroquial da Igreja de Santa Madre Cabrini, desenvolvendo atividades de socialização, recreação, trabalhos manuais e recebendo complementação alimentar. Em 1987, sensibilizados pela campanha da fraternidade que tinha como Tema: Fraternidade e o Menor e como lema: “Quem acolhe o menor a mim acolhe”, assumimos definitivamente o trabalho assistencial àquelas crianças. Nesse mesmo ano foi regularizada a documentação necessária para o pleno funcionamento da entidade. Com o decorrer dos anos, o trabalho foi conhecido por diretores da empresa A.W. Faber Castell S. A., que sensibilizados com a problemática da criança e do adolescente, possibilitaram a edificação do prédio (465,26m²). A construção iniciou em 1989 e foi inaugurada em 1990. Atualmente, a área construída é de 687,70m² e tem sido utilizada para vários projetos sociais, abrangendo não só o público infantil como também a população adulta da comunidade.</p>
      </article>

      <article class="content">
        <h2 class="title">Nossa missão</h2>
        <p class="paragraph">A missão da entidade é proporcionar serviço social, apoio educacional, complementação alimentar, assistência moral, religiosa, cultural, emocional e recreativa às crianças/adolescentes sem distinção: de credo, religião, política, cor ou raça e motivar a transformação e valorização da vida.</p>
      </article>
    </section>
  </section>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .HomeScreen
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
    .HomeScreen
      flex-direction: row
      height: 900px

      > .aside
        width: 50%

      > .HomeContent
        width: 50%

  @media screen and (min-width: 1200px)
    .HomeScreen
      height: 900px

      > .aside
        width: 30%

      > .HomeContent
        width: 70%

  .HomeContent
    height: 100%

    > .content
      padding: 24px
      background-size: cover
      background-position: center

      > .title
        { $typography-title }
        font-size: 3rem
        color: $color-contrast

      > .paragraph
        margin-top: 12px
        font-size: 1.6rem
        line-height: 1.4
        color: #FFFFFF

    > .content:nth-child(1)
      background-image: url('~@/assets/images/Content-1.png')

    > .content:nth-child(2)
      background-image: url('~@/assets/images/Content-3.png')

    > .content:nth-child(3)
      background-image: url('~@/assets/images/Content-2.png')


  @media screen and (min-width: 1200px)
    .HomeContent > .content
      float: left
      width: 50%
      min-height: 50%

    .HomeContent > .content:nth-child(2)
      float: right
      height: 100%
</style>
