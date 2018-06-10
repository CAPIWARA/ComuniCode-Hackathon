<template>
  <section class="CheckoutScreen">
    <impact-form class="half-section">
      <impact-entry
        v-model="address.street"
        label="Logradouro"
        name="Logradouro"
      />
      <impact-entry
        v-model="address.number"
        label="Número"
        name="Número"
      />
      <impact-entry
        v-model="address.state"
        label="Estado"
        name="Estado"
      />
      <impact-entry
        v-model="address.city"
        label="Cidade"
        name="Cidade"
      />
      <impact-entry
        v-model="address.code"
        label="CEP"
        name="CEP"
      />
    </impact-form>

    <impact-form class="half-section">
      <impact-entry
        v-model="card.number"
        label="Número do Cartão"
        name="Número do Cartão"
      />
      <impact-entry
        v-model="card.name"
        label="Nome"
        name="Nome"
      />
      <impact-entry
        v-model="card.date"
        label="Data de Vencimento"
        name="Data de Vencimento"
      />
      <impact-entry
        v-model="card.code"
        label="Código de Segurança"
        name="Código de Segurança"
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
      align-items: center
</style>
