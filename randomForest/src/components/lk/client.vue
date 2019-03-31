<template>
  <v-container>

    <v-layout v-if="loading">
      <v-flex xs12 class="text-xs-center pt-5">
        <v-progress-circular indeterminate :size="100" :width="4" color="purple">
        </v-progress-circular>
      </v-flex>
    </v-layout>

    <v-layout v-else-if="!loading">
      <v-flex md12 class="text-xs-center">
        <h1>Сводна информация </h1>

      </v-flex>
    </v-layout>
    <v-layout justify-space-around>
      <v-flex md6 class="mt-5">
        <line-example v-if="loaded" :chart-data="value" :chart-labels="label"></line-example>
        <wind-example v-if="loaded" :chart-data="valueHumidity" :chart-labels="labelHumidity"></wind-example>
        <humidity-example v-if="loaded" :chart-data="valueWind" :chart-labels="labelWind"></humidity-example>
      </v-flex>
      <v-flex md6 class="mt-5 ml-5">
        <h1>Сводка</h1>
      </v-flex>
    </v-layout>

  </v-container>
</template>

<script>
  export default {
    name: "client",

    data() {
      return {
        loaded: false,
        label: [],
        value: [],
        labelHumidity: [],
        valueHumidity: [],
        labelWind: [],
        valueWind: [],
      }
    },
    computed: {
      loading() {
        return this.$store.getters.loading
      },
      getTemps() {
        let obj = this.$store.getters.getTempJson;
        if (obj) {
          this.label = obj['Date'];
          this.value = obj['Value']
        }
      },
      getHumiditys() {
        let obj = this.$store.getters.getHumidityJson;
        if (obj) {
          this.labelHumidity = obj['Date'];
          this.valueHumidity = obj['Value']
        }
      },
      getWinds() {
        let obj = this.$store.getters.getWindJson;
        if (obj) {
          this.labelWind = obj['Date'];
          this.valueWind = obj['Value']
        }
      }
    },
    created() {
      this.loaded = false
      this.$store.dispatch('getTemp').then(() => {
        this.getTemps
      })
      this.$store.dispatch('getWind').then(() => {
        this.getWinds
      })
      this.$store.dispatch('getHumidity').then(() => {
        this.getHumiditys
        this.loaded = true
      })
    },


  }
</script>

<style>
  .small {
    max-width: 600px;
    margin: 150px auto;
  }

</style>
