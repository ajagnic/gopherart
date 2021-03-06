<template>
  <div>
    <v-progress-linear v-model="progress" :active="loading" rounded />
    <v-card>
      <div v-if="dataURL == null">
        <v-card-title>Generate Image</v-card-title>
        <v-card-subtitle>
          Upload a JPEG or PNG image below to begin
        </v-card-subtitle>
        <v-container>
          <image-upload @file-upload="input = $event" />
        </v-container>
        <v-card-actions>
          <v-btn
            x-large
            block
            plain
            :disabled="input === null"
            :loading="loading"
            @click="processFile"
          >
            Generate
          </v-btn>
        </v-card-actions>
      </div>
    </v-card>
    <v-fade-transition hide-on-leave>
      <v-card v-show="dataURL != null">
        <v-container>
          <v-btn
            class="top-fab"
            small
            fab
            absolute
            top
            left
            color="primary"
            @click="processFile"
          >
            <v-icon color="black">mdi-refresh</v-icon>
          </v-btn>
          <image-controls :values.sync="params" />
          <a :href="dataURL" :download="filename">
            <v-btn
              id="download-btn"
              class="top-fab"
              small
              fab
              absolute
              top
              right
            >
              <v-icon>mdi-download</v-icon>
            </v-btn>
          </a>
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
          <div
            class="image"
            :style="{ backgroundImage: 'url(' + dataURL + ')' }"
          />
        </v-container>
      </v-card>
    </v-fade-transition>
  </div>
</template>

<script>
export default {
  data: () => ({
    input: null,
    dataURL: null,
    filename: null,
    enc: null,
    worker: null,
    loading: false,
    progress: 0,
    params: {
      iterations: 10000,
      polygonSidesMin: 3,
      polygonSidesMax: 5,
      polygonFill: 1.0,
      polygonColor: 0.0,
      polygonSizeRatio: 0.1,
      pixelShake: 0.0,
      pixelSpin: 0,
      newWidth: 0,
      newHeight: 0,
      greyscale: false,
      invertScaling: false,
    },
  }),

  mounted() {
    this.worker = new Worker('js/go_worker.js')
    this.worker.addEventListener('message', this.receiveMessage)
  },

  methods: {
    processFile() {
      this.loading = true
      this.filename = 'gopherart-' + this.input.name
      this.enc = this.input.type
      this.input.arrayBuffer().then((blob) => {
        const bytes = new Uint8Array(blob)
        const size = bytes.byteLength
        const paramsStr = JSON.stringify(this.params)
        this.worker.postMessage({ b: bytes, len: size, opt: paramsStr })
      })
    },

    receiveMessage(e) {
      const msg = e.data
      if (msg.type === 'progress') {
        this.progress = msg.value
      } else {
        this.loading = false
        this.progress = 0
        this.dataURL = 'data:' + this.enc + ';base64,' + JSON.parse(msg.value)
        this.$emit('image-loaded', true)
      }
    },

    closeProcessing() {
      this.input = null
      this.dataURL = null
      this.$emit('image-close', true)
    },
  },
}
</script>

<style>
div.image {
  height: 500px;
  max-height: 80vh;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

.top-fab {
  margin-top: 30px;
}

#close-btn {
  right: 15px;
}

#ctrl-btn {
  margin-left: 50px;
}

#download-btn {
  right: 65px;
}
</style>
