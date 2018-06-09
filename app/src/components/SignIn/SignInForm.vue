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

    <impact-button>Entrar</impact-button>
  </form>
</template>

<script>
  import FormEntry from '@/components/Form/FormEntry';
  import ImpactButton from '@/components/Impact/ImpactButton';

  export default {
    components: { FormEntry, ImpactButton },
    data () {
      return {
        email: '',
        password: ''
      };
    }
  };
</script>
