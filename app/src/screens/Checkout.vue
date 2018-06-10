<template>
  <section class="CheckoutScreen">
    <h3>Valor fíxo de R$ 30,00.</h3>

    <impact-form class="half-section">
      <impact-entry
        v-modal="address.street"
        label="Logradouro"
      />
      <impact-entry
        v-modal="address.number"
        label="Número"
      />
      <impact-entry
        v-modal="address.state"
        label="Estado"
      />
      <impact-entry
        v-modal="address.city"
        label="Cidade"
      />
      <impact-entry
        v-modal="address.code"
        label="CEP"
      />
    </impact-form>

    <impact-form class="half-section">
      <impact-entry
        v-modal="card.number"
        label="Número do Cartão"
      />
      <impact-entry
        v-modal="card.name"
        label="Nome"
      />
      <impact-entry
        v-modal="card.date"
        label="Data de Vencimento"
      />
      <impact-entry
        v-modal="card.code"
        label="Código de Segurança"
      />
    </impact-form>
    <impact-form class="full-section">
      <impact-button size="xl" @click="donate()">Confirmar Doação</impact-button>
    </impact-form>
  </section>
</template>

<script>
  import ImpactForm from '@/components/Impact/ImpactForm';
  import ImpactEntry from '@/components/Impact/ImpactEntry';
  import ImpactButton from '@/components/Impact/ImpactButton';
  import { mapActions } from 'vuex';
  import { CHECKOUT } from '@/store/types';

  export default {
    components: { ImpactForm, ImpactEntry, ImpactButton },
    data: () => ({
      address: {
        street: '',
        number: '',
        state: '',
        city: '',
        code: ''
      },
      card: {
        number: '',
        name: '',
        date: '',
        code: ''
      }
    }),
    methods: mapActions({
      donate: CHECKOUT
    })
  };
</script>

<style lang="stylus">
  .CheckoutScreen
    display: flex
    flex-wrap: wrap

    > .half-section
      width: 50%

    > .full-section
      width: 100%
</style>
