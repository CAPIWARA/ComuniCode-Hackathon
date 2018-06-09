<template>
  <fieldset>
    <label>{{ label }}</label>
    <input
      v-bind="$attrs"
      :name="name"
      :value="value"
      @input="$emit('input', $event.target.value)"
    />
    <span v-if="!!error">{{ error }}</span>
  </fieldset>
</template>

<script>
  export default {
    props: {
      name: {
        type: String,
        required: true
      },
      value: {
        type: String,
        required: true
      },
      error: {
        type: String,
        default: null
      },
      label: {
        type: String,
        default: null
      }
    }
  };
</script>
