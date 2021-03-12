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
      <v-btn x-large block plain @click="processFile">Submit</v-btn>
    </v-card-actions>
    <img v-if="dataURL != null" :src="dataURL" />
  </v-card>
</template>

<script>
export default {
  data: () => ({
    input: null,
    dataURL: null,
  }),

  methods: {
    processFile() {
      this.input.arrayBuffer().then((blob) => {
        const imgObj = new Int8Array(blob)
        const img = Object.values(imgObj)
        const img64 = window.processImage(JSON.stringify(img))
        this.dataURL = 'data:image/jpeg;base64,' + JSON.parse(img64)
      })
    },
  },
}
</script>
