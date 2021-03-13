<template>
  <v-card>
    <div v-if="dataURL == null">
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
    </div>
    <div v-if="dataURL != null">
      <v-card-actions>
        <image-controls @update:params="params = $event" />
        <v-spacer />
        <v-btn @click="processFile">Generate</v-btn>
      </v-card-actions>
      <v-card-text>
        <v-row>
          <v-spacer />
          <img width="90%" :src="dataURL" />
          <v-spacer />
        </v-row>
      </v-card-text>
    </div>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    input: null,
    dataURL: null,
    params: {
      iterations: 10000,
      width: 600,
      height: 600,
      polygonSidesMin: 3,
      polygonSidesMax: 5,
      polygonFillChance: 1.0,
      polygonColorChance: 0.0,
      polygonSizeRatio: 0.1,
      pixelShake: 0.0,
      greyscale: false,
    },
  }),

  methods: {
    processFile() {
      this.input.arrayBuffer().then((blob) => {
        const imgObj = new Int8Array(blob)
        const img = Object.values(imgObj)
        const imgStr = JSON.stringify(img)
        const paramsStr = JSON.stringify(this.params)
        const img64Str = window.processImage(imgStr, paramsStr)
        this.dataURL = 'data:image/jpeg;base64,' + JSON.parse(img64Str)
      })
    },
  },
}
</script>
