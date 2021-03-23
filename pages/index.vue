<template>
  <div>
    <br />
    <v-card id="main-card" class="mx-auto" max-width="600">
      <div v-if="!imageLoaded">
        <div v-if="!showProcessing">
          <v-card-text>
            <v-row>
              <v-spacer />
              <div id="scaler">
                <div id="spinner"></div>
              </div>
              <v-spacer />
            </v-row>
          </v-card-text>
          <v-card-text class="text-center">
            <p>
              Welcome to gopherart.dev!<br />
              Perform artistic image-processing<br />
              in your browser, using<br />
              Go and Web Assembly.
              <!-- This site uses Go WebAssembly<br />
              to perform image processing in your browser. -->
            </p>
          </v-card-text>
        </div>
        <v-card-actions v-if="!showProcessing">
          <v-btn
            large
            block
            text
            outlined
            color="primary"
            @click="showProcessing = true"
          >
            Upload Image
          </v-btn>
        </v-card-actions>
      </div>
      <v-fade-transition>
        <image-processing
          v-show="showProcessing"
          @image-loaded="imageLoaded = true"
          @image-close="imageLoaded = false"
        />
      </v-fade-transition>
    </v-card>
  </div>
</template>

<script>
export default {
  data: () => ({
    showProcessing: false,
    imageLoaded: false,
  }),

  head() {
    return {
      title: 'home',
      script: [{ src: 'js/instantiate_go.js' }],
    }
  },
}
</script>

<style scoped>
#main-card {
  background: transparent;
}

#spinner {
  width: 150px;
  height: 150px;
  background: url('~assets/logo.png');
  background-size: contain;
  -webkit-animation: rotation 40s infinite linear;
  -moz-animation: rotation 40s infinite linear;
  -o-animation: rotation 40s infinite linear;
  animation: rotation 40s infinite linear;
}

#scaler {
  width: 150px;
  height: 150px;
  -webkit-animation: scale 5s infinite;
  -moz-animation: scale 5s infinite;
  -o-animation: scale 5s infinite;
  animation: scale 5s infinite;
}

@keyframes rotation {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(359deg);
  }
}

@keyframes scale {
  0%,
  100% {
    transform: scale(100%);
  }

  50% {
    transform: scale(110%);
  }
}
</style>
