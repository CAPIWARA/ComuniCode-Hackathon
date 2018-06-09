<template>
  <form novalidate @submit.prevent="submit()">
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

    <button type="submit">Entrar</button>
  </form>
</template>

<script>
  import FormEntry from '@/components/Form/FormEntry';

  export default {
    components: { FormEntry },
    data () {
      return {
        email: '',
        password: ''
      };
    }
  };
</script>
