<template>
  <v-menu :close-on-content-click="false">
    <template #activator="{ on, attrs }">
      <v-btn v-bind="attrs" v-on="on">Advanced</v-btn>
    </template>
    <v-card min-width="300" @mouseleave="$emit('update:params', params)">
      <v-card-text>
        <v-text-field
          v-model.number="params.iterations"
          label="Iterations"
          type="number"
        />
        <v-text-field
          v-model.number="params.width"
          label="Width"
          type="number"
        />
        <v-text-field
          v-model.number="params.height"
          label="Height"
          type="number"
        />
        <v-range-slider
          label="PolygonSides"
          thumb-label
          min="2"
          max="20"
          :value="sides"
          @change="setPolygonSides"
        />
        <v-slider
          v-model="params.polygonFillChance"
          label="PolygonFill"
          thumb-label
          :min="min"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.polygonColorChance"
          label="PolygonColor"
          thumb-label
          :min="min"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.polygonSizeRatio"
          label="PolygonSize"
          thumb-label
          :min="min"
          :max="max"
          :step="step"
        />
        <v-slider
          v-model="params.pixelShake"
          label="PixelShake"
          thumb-label
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
  data: () => ({
    min: 0.0,
    max: 1.0,
    step: 0.1,
    sides: [3, 5],
    params: {
      iterations: 10000,
      width: 600,
      height: 600,
      polygonSidesMin: 3,
      polygonSidesMax: 5,
      polygonFillChance: 1.0,
      polygonColorChance: 0.0,
      polygonSizeRatio: 0.1,
      pixelShake: 0.0,
      greyscale: false,
    },
  }),

  methods: {
    setPolygonSides(val) {
      this.params.polygonSidesMin = val[0]
      this.params.polygonSidesMax = val[1]
    },
  },
}
</script>
