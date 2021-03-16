<template>
  <v-card
    id="dropper"
    flat
    height="300"
    @drop.prevent="onDrop($event)"
    @dragover.prevent
    @dragleave.prevent
  >
    <v-card-text class="text-center">
      <v-icon size="100" :color="iconColor" :style="iconStyle">
        mdi-cloud-upload-outline
      </v-icon>
    </v-card-text>
    <v-card-title v-if="!fileDropped" class="justify-center">
      Drop a file, or&nbsp;
      <label>
        <v-file-input v-show="false" v-model="file" @change="onSelect" />
        <a>select one.</a>
      </label>
    </v-card-title>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    file: null,
    fileDropped: false,
    iconColor: 'white',
    iconStyle: {
      'padding-top': '30px',
    },
  }),

  methods: {
    onDrop(e) {
      this.fileDropped = true
      this.iconColor = 'primary'
      this.iconStyle = {}
      this.$emit('file-drop', e)
    },

    onSelect() {
      this.fileDropped = true
      this.iconColor = 'primary'
      this.iconStyle = {}
      this.$emit('file-select', this.file)
    },
  },
}
</script>

<style scoped>
#dropper {
  outline-style: dashed;
}
</style>
