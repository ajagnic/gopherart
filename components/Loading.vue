<template>
  <v-overlay v-if="loading" opacity="0.6">
    <p color="primary">Loading...</p>
  </v-overlay>
</template>

<script>
export default {
  data: () => ({
    loading: false,
  }),

  methods: {
    start() {
      this.loading = true
    },

    finish() {
      this.loading = false
    },
  },
}
</script>
