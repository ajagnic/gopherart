<template>
  <v-card class="mx-auto" max-width="700">
    <v-card-text>
      <v-row>
        <v-spacer />
        <img src="@/assets/gopher.png" width="30%" />
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
    <image-processing />
  </v-card>
</template>

<script>
export default {
  head() {
    return {
      script: [{ src: 'js/instantiate_go.js' }],
    }
  },
}
</script>
