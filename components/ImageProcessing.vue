<template>
  <v-card>
    <v-card-title>Generate Image</v-card-title>
    <v-card-subtitle>Upload your own image to process</v-card-subtitle>
    <v-card-text>
      <v-file-input
        v-model="input"
        hint="*.jpeg, *.png"
        persistent-hint
        accept="image/jpeg, image/png"
      />
    </v-card-text>
    <v-card-actions>
      <v-btn x-large block plain @click="test">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    input: null,
  }),

  methods: {
    test() {
      this.input.arrayBuffer().then((result) => {
        const arr = new Int8Array(result)
        console.log(arr)
      })
    },
  },
}
</script>
