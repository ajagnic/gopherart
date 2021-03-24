<template>
  <v-menu
    v-model="menu"
    transition="scale-transition"
    max-height="75%"
    :close-on-content-click="false"
  >
    <template #activator="{ on, attrs }">
      <v-btn
        id="ctrl-btn"
        class="top-fab"
        small
        fab
        absolute
        top
        left
        v-bind="attrs"
        v-on="on"
      >
        <v-icon>mdi-tune</v-icon>
      </v-btn>
    </template>
    <v-card max-width="300" @mouseleave="updateParams">
      <v-card-text>
        <v-text-field
          v-model.number="params.iterations"
          suffix="iterations"
          hint="High iterations will take longer to process"
          outlined
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
        <v-btn v-if="!showWH" block @click="showWH = true">
          Set Width/Height
        </v-btn>
        <div v-else>
          <v-row>
            <v-col>
              <v-text-field
                v-model.number="params.width"
                prefix="Width:"
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.height"
                prefix="Height:"
                outlined
              />
            </v-col>
          </v-row>
        </div>
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
    showWH: false,
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
      this.showWH = false
      this.$emit('update:params', this.params)
    },
  },
}
</script>
