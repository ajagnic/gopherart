<template>
  <div>
    <v-app-bar :collapse="collapse" dense app>
      <!-- <v-app-bar-nav-icon @click="drawer = !drawer" /> -->
      <v-btn icon @click="collapse = !collapse">
        <v-icon>mdi-chevron-left</v-icon>
      </v-btn>
      <v-spacer />
      <v-app-bar-title>
        gopherart.<span class="primary--text text--accent-3">dev</span>
      </v-app-bar-title>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" app>
      <v-list nav>
        <v-list-item
          v-for="page in pages"
          :key="page.title"
          nuxt
          :to="page.link"
        >
          <v-list-item-icon>
            <v-icon v-text="page.icon" />
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title v-text="page.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  data: () => ({
    collapse: false,
    drawer: false,
    pages: [{ title: 'Home', icon: 'mdi-home', link: '/' }],
  }),
}
</script>
