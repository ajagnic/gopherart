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
        <image-controls :values.sync="params" />
        <v-spacer />
        <a :href="dataURL" :download="filename">
          <v-btn>Download</v-btn>
        </a>
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
    filename: null,
    params: {
      iterations: 10000,
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
      this.filename = 'gopherart-' + this.input.name
      const enc = this.input.type
      this.input.arrayBuffer().then((blob) => {
        const bytes = new Uint8Array(blob)
        const size = bytes.byteLength
        const paramsStr = JSON.stringify(this.params)
        const img64Str = window.processImage(bytes, size, paramsStr)
        this.dataURL = 'data:' + enc + ';base64,' + JSON.parse(img64Str)
      })
    },
  },
}
</script>
