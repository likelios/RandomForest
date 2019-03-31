<template>
  <v-container>

    <v-layout v-if="loading">
      <v-flex xs12 class="text-xs-center pt-5">
        <v-progress-circular indeterminate :size="100" :width="4" color="purple">
        </v-progress-circular>
      </v-flex>
    </v-layout>

    <v-layout v-else-if="!loading">
      <!--    <v-layout >-->
      <v-flex md12 class="text-xs-center">
        <h1>Сводна информация </h1>
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex md12 class="mt-5">
        <line-example v-if="loaded" :chart-data="value" :chart-labels="label"></line-example>
        <wind-example></wind-example>
        <humidity-example></humidity-example>
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
        value: []
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

      }
    },
    created() {
      this.$store.dispatch('getTemp').then(() => {
        this.getTemps
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
