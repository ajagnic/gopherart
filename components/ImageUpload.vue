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
        <v-file-input
          v-show="false"
          v-model="file"
          accept="image/jpeg, image/png"
          @change="onUpload"
        />
        <a>select one.</a>
      </label>
    </v-card-title>
    <v-alert
      v-if="error"
      class="mx-auto"
      type="error"
      max-width="400"
      dismissible
    >
      {{ status }}
    </v-alert>
    <v-card-title
      v-if="fileDropped && !error"
      class="justify-center font-italic"
    >
      {{ filename }}
    </v-card-title>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    error: false,
    status: null,
    file: null,
    filename: null,
    fileDropped: false,
    iconColor: 'white',
    iconStyle: {
      'padding-top': '30px',
    },
  }),

  methods: {
    showError(status) {
      this.file = null
      this.fileDropped = false
      this.error = true
      this.iconColor = 'error'
      this.iconStyle = {
        'padding-top': '30px',
      }
      this.status = status
    },

    showSuccess() {
      this.fileDropped = true
      this.error = false
      this.iconColor = 'primary'
      this.iconStyle = {}
      this.filename = this.file.name
    },

    onDrop(e) {
      const fileList = e.dataTransfer.files
      this.file = e.dataTransfer.files[0]
      if (fileList.length > 1) {
        this.showError('Please select a single file.')
      }
      this.onUpload()
    },

    onUpload() {
      if (this.file != null) {
        const type = this.file.type
        if (type !== 'image/jpeg' && type !== 'image/png') {
          this.showError('Please select a JPEG or PNG image.')
        } else {
          this.showSuccess()
        }
      }
      this.$emit('file-upload', this.file)
    },
  },
}
</script>

<style scoped>
#dropper {
  outline-style: dashed;
  background: rgba(50, 50, 50, 0.5);
}
</style>
