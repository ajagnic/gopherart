<template>
  <v-menu
    v-model="menu"
    transition="scale-transition"
    :close-on-content-click="false"
  >
    <template #activator="{ on, attrs }">
      <v-btn class="top-fab" medium fab absolute left v-bind="attrs" v-on="on">
        <v-icon>mdi-tune-variant</v-icon>
      </v-btn>
    </template>
    <v-card min-width="300" @mouseleave="updateParams">
      <v-card-text>
        <v-slider
          v-model="params.iterations"
          label="Iterations"
          min="1"
          max="100000"
          step="1000"
        />
        <v-range-slider
          label="Sides"
          thumb-label
          min="2"
          max="20"
          :value="sides"
          @change="setPolygonSides"
        />
        <v-slider
          v-model="params.polygonFillChance"
          label="Fill"
          :min="min"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.polygonColorChance"
          label="Color"
          :min="min"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.polygonSizeRatio"
          label="Size"
          min="0.01"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.pixelShake"
          label="Shake"
          :min="min"
          :max="max"
          :step="step"
        />
        <v-switch v-model="params.greyscale" label="Greyscale" />
      </v-card-text>
    </v-card>
  </v-menu>
</template>

<script>
export default {
  props: {
    values: {
      type: Object,
      default: () => ({}),
    },
  },

  data: () => ({
    menu: false,
    min: 0.0,
    max: 1.0,
    step: 0.05,
    sides: [3, 5],
    params: {},
  }),

  mounted() {
    this.params = this.values
  },

  methods: {
    setPolygonSides(val) {
      this.params.polygonSidesMin = val[0]
      this.params.polygonSidesMax = val[1]
    },

    updateParams() {
      this.menu = false
      this.$emit('update:params', this.params)
    },
  },
}
</script>
