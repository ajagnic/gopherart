<template>
  <v-card>
    <div v-if="dataURL == null">
      <v-container>
        <image-upload @file-upload="input = $event" />
      </v-container>
      <v-card-actions>
        <v-btn
          x-large
          block
          plain
          :disabled="input === null"
          @click="processFile"
        >
          Generate
        </v-btn>
      </v-card-actions>
    </div>
    <div v-if="dataURL != null">
      <v-container>
        <v-btn id="close" small fab absolute top right @click="closeProcessing">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-img :src="dataURL">
          <image-controls :values.sync="params" />
          <v-btn class="top-fab" small fab absolute right @click="processFile">
            <v-icon>mdi-refresh</v-icon>
          </v-btn>
          <a :href="dataURL" :download="filename">
            <v-btn
              class="bottom-fab"
              small
              fab
              absolute
              bottom
              right
              color="primary"
            >
              <v-icon color="black">mdi-download</v-icon>
            </v-btn>
          </a>
        </v-img>
      </v-container>
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
        this.$emit('image-loaded', true)
      })
    },

    closeProcessing() {
      this.dataURL = null
      this.$emit('image-close', true)
    },
  },
}
</script>

<style>
.bottom-fab {
  margin-bottom: 50px;
}

.top-fab {
  margin-top: 10px;
}

#close {
  right: -10px;
  margin-top: 10px;
}
</style>
