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
        <v-btn class="top-fab" small fab absolute top left @click="processFile">
          <v-icon>mdi-refresh</v-icon>
        </v-btn>
        <image-controls :values.sync="params" />
        <v-btn
          id="close-btn"
          class="top-fab"
          small
          fab
          absolute
          top
          right
          @click="closeProcessing"
        >
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <a :href="dataURL" :download="filename">
          <v-btn
            id="download-btn"
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
        <img width="100%" height="100%" :src="dataURL" />
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
.top-fab {
  margin-top: 10px;
}

#close-btn {
  right: -10px;
}

#ctrl-btn {
  margin-left: 50px;
}

#download-btn {
  margin-bottom: 10px;
  right: -10px;
}
</style>
