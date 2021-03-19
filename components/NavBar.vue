<template>
  <div>
    <v-app-bar :shrink-on-scroll="shrink" dense app>
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-btn icon plain @click="shrinkBar">
        <v-icon>{{ shrinkIcon }}</v-icon>
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
    drawer: false,
    shrink: true,
    shrinkIcon: 'mdi-chevron-up',
    pages: [{ title: 'Home', icon: 'mdi-home', link: '/' }],
  }),

  methods: {
    shrinkBar() {
      this.shrink = !this.shrink
      if (this.shrink === true) {
        this.shrinkIcon = 'mdi-chevron-up'
      } else {
        this.shrinkIcon = 'mdi-chevron-down'
      }
    },
  },
}
</script>
