<template>
  <v-card>
    <img :src="'data:image/jpeg;base64,' + dataURL" />
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
    dataURL: null,
  }),

  methods: {
    test() {
      this.input.arrayBuffer().then((blob) => {
        const arrObj = new Int8Array(blob)
        const arr = Object.values(arrObj)
        const jsonArr = JSON.stringify(arr)
        const jb = window.processImage(jsonArr)
        this.dataURL = JSON.parse(jb)
      })
    },
  },
}
</script>
