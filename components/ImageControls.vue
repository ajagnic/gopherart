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
    <v-card max-width="300" color="#333333" @mouseleave="updateParams">
      <v-card-text>
        <v-text-field
          v-model.number="params.iterations"
          suffix="iterations"
          hint="High iteration will take longer to process"
          outlined
        />
        <div v-if="!showValues">
          <v-range-slider
            label="Sides"
            thumb-label
            min="2"
            max="20"
            :value="sides"
            @change="setPolygonSides"
          />
          <v-slider
            v-model="params.polygonFill"
            label="Fill"
            :min="min"
            :max="max"
            :step="step"
          />
          <v-slider
            v-model="params.polygonColor"
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
          <v-slider
            v-model="params.pixelSpin"
            label="Spin"
            min="0"
            max="360"
            step="10"
          />
        </div>
        <div v-else>
          <v-row>
            <v-col>
              <v-text-field
                v-model.number="params.polygonSidesMin"
                dense
                label="Min Sides"
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.polygonSidesMax"
                dense
                label="Max Sides"
                outlined
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                v-model.number="params.polygonFill"
                dense
                label="Fill"
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.polygonColor"
                dense
                label="Color"
                outlined
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                v-model.number="params.newWidth"
                label="Width"
                dense
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.newHeight"
                label="Height"
                dense
                outlined
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                v-model.number="params.polygonSizeRatio"
                dense
                label="Size"
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.pixelShake"
                dense
                label="Shake"
                outlined
              />
            </v-col>
            <v-col>
              <v-text-field
                v-model.number="params.pixelSpin"
                dense
                label="Spin"
                outlined
              />
            </v-col>
          </v-row>
        </div>
        <v-row justify="space-around">
          <v-switch v-model="params.greyscale" label="Greyscale" />
          <v-switch v-model="params.invertScaling" label="Invert Scaling" />
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn icon @click="showAdvanced">
          <v-icon>{{ advIcon }}</v-icon>
        </v-btn>
      </v-card-actions>
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
    showValues: false,
    advIcon: 'mdi-cog',
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

    showAdvanced() {
      if (!this.showValues) {
        this.showValues = true
        this.advIcon = 'mdi-cog-off'
      } else {
        this.showValues = false
        this.advIcon = 'mdi-cog'
      }
    },
  },
}
</script>

<style scoped>
.row {
  margin-bottom: -20px;
}
</style>
