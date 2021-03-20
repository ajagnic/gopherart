<template>
  <div>
    <v-card class="mx-auto" max-width="600">
      <div v-if="!imageLoaded">
        <v-card-text>
          <v-row>
            <v-spacer />
            <div id="scaler">
              <div id="logo"></div>
            </div>
            <v-spacer />
          </v-row>
        </v-card-text>
        <v-card-text>
          <v-row>
            <v-spacer />
            <p>Welcome to gopherart.dev!<br /></p>
            <v-spacer />
          </v-row>
        </v-card-text>
      </div>
      <image-processing @image-loaded="imageLoaded = true" />
    </v-card>
  </div>
</template>

<script>
export default {
  data: () => ({
    imageLoaded: false,
  }),

  head() {
    return {
      script: [{ src: 'js/instantiate_go.js' }],
    }
  },
}
</script>

<style scoped>
#logo {
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
