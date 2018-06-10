<template>
  <impact-form @submit="submit()">
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
      v-model="password"
      v-validate="'required|min:8|max:16'"
      name="senha"
      type="password"
      label="Senha"
      autocomplete="current-password"
      :error="errors.first('senha')"
    />

    <impact-button>Entrar</impact-button>
  </impact-form>
</template>

<script>
  import ImpactForm from '@/components/Impact/ImpactForm';
  import ImpactEntry from '@/components/Impact/ImpactEntry';
  import ImpactButton from '@/components/Impact/ImpactButton';

  export default {
    components: { ImpactForm, ImpactEntry, ImpactButton },
    data () {
      return {
        email: '',
        password: ''
      };
    },
    methods: {
      async submit () {
        const isValid = await this.$validator.validateAll();

        if (!isValid)
          return

        this.$emit('submit', {
          email: this.email,
          password: this.password
        });
      }
    }
  };
</script>
