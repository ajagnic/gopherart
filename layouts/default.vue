<template>
  <v-app>
    <v-app-bar shrink-on-scroll prominent app>
      <v-spacer />
      <v-app-bar-title>gopherart.dev</v-app-bar-title>
    </v-app-bar>
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
  background-image: url('~assets/background.jpg');
  height: 100vh;
}
</style>
