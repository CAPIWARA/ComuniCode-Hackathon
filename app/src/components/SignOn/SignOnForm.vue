<template>
  <form novalidate @submit.prevent="submit()">
    <form-entry
      v-model="name"
      v-validate="'required'"
      name="nome"
      type="text"
      label="Nome"
      autocomplete="name"
      :error="errors.first('nome')"
    />

    <form-entry
      v-model="email"
      v-validate="'required|email'"
      name="e-mail"
      type="email"
      label="E-Mail"
      autocomplete="email"
      :error="errors.first('e-mail')"
    />

    <form-entry
      v-model="password"
      v-validate="'required|min:8|max:16'"
      name="senha"
      type="password"
      label="Senha"
      autocomplete="current-password"
      :error="errors.first('senha')"
    />

    <form-entry
      v-model="passwordConfirmation"
      v-validate="'required|min:8|max:16'"
      name="confirmação da senha"
      type="password"
      label="Confirmação da Senha"
      autofill="off"
      autocomplete="off"
      :error="errors.first('confirmação da senha')"
    />

    <button type="submit">Entrar</button>
  </form>
</template>

<script>
  import FormEntry from '@/components/Form/FormEntry';

  export default {
    components: { FormEntry },
    data () {
      return {
        name: '',
        email: '',
        password: '',
        passwordConfirmation: '',
        document: ''
      };
    },
    methods: {
      async submit () {
        this.fields.validateAll()
      }
    }
  };
</script>
