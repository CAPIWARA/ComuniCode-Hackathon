<template>
  <impact-form novalidate @submit.prevent="submit()">
    <impact-entry
      v-model="name"
      v-validate="'required'"
      name="nome"
      type="text"
      label="Nome"
      autocomplete="name"
      :error="errors.first('nome')"
    />

    <impact-entry
      v-model="email"
      v-validate="'required|email'"
      name="e-mail"
      type="email"
      label="E-Mail"
      autocomplete="email"
      :error="errors.first('e-mail')"
    />

    <impact-entry
      v-validate="'required|min:8|max:16'"
      name="senha"
      type="password"
      label="Senha"
      autocomplete="current-password"
      :value="password"
      :error="errors.first('senha')"
      @input="onInputPassword($event)"
    />

    <impact-entry
      v-model="confirmation"
      v-validate="'required|min:8|max:16|equals:password'"
      name="confirmacao"
      type="password"
      label="Confirmação da Senha"
      autofill="off"
      autocomplete="off"
      :error="errors.first('confirmacao')"
    />

    <div class="button-group">
      <impact-button
        type="submit"
        isAlternative
      >
        Já possui
        <br>uma conta?
      </impact-button>

      <impact-button type="submit">
        Cadastrar-se<br>
        como doador
      </impact-button>
    </div>
  </impact-form>
</template>

<script>
  import ImpactForm from '@/components/Impact/ImpactForm';
  import ImpactEntry from '@/components/Impact/ImpactEntry';
  import ImpactButton from '@/components/Impact/ImpactButton';
  import { Validator } from 'vee-validate';

  export default {
    components: { ImpactForm, ImpactEntry, ImpactButton },
    data () {
      return {
        name: '',
        email: '',
        password: '',
        confirmation: '',
        document: ''
      };
    },
    methods: {
      async submit () {
        const form = {
          name: this.name,
          email: this.email,
          document: this.document,
          password: this.password,
          confirmation: this.confirmation,
        };
        const isValid = await this.$validator.validateAll();
        if (!isValid)
          return;
        this.$emit('submit', form);
      },
      onInputPassword (value) {
        this.password = value
        this.$validator.validate('confirmacao')
      }
    },
    created () {
      Validator.extend('equals', {
        getMessage: (name) => console.log(name) || 'Senha e confirmação não coincide.',
        validate: (value, properties) => {
          const isValid = properties.every((property) => value === this[property]);
          return isValid;
        }
      });
    }
  };
</script>

<style lang="stylus">
  .button-group
    display: flex
    width: 100%
    margin-top: 24px

    > .ImpactButton
      margin-left: auto

    > .ImpactButton + .ImpactButton
      margin-left: 12px
</style>

