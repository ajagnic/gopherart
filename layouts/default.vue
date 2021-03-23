<template>
  <v-app>
    <nav-bar />
    <v-main>
      <v-container fluid>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <v-spacer />
      {{ new Date().getFullYear() }} - Adrian Agnic
      <v-spacer />
    </v-footer>
  </v-app>
</template>

<style scoped>
.container {
  background-image: url('~assets/background.jpeg');
  height: 100%;
  width: 100vw;
}
</style>
