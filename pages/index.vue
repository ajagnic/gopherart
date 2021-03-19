<template>
  <div>
    <v-card class="mx-auto" max-width="700">
      <div v-if="!imageLoaded">
        <v-card-text>
          <v-row>
            <v-spacer />
            <img src="@/assets/logo.png" width="20%" />
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
