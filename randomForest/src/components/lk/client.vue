<template>
  <v-container>

<!--    <v-layout v-if="loading">-->
<!--      <v-flex xs12 class="text-xs-center pt-5">-->
<!--        <v-progress-circular indeterminate :size="100" :width="4" color="purple">-->
<!--        </v-progress-circular>-->
<!--      </v-flex>-->
<!--    </v-layout>-->

<!--    <v-layout v-else-if="!loading">-->
    <v-layout >
      <v-flex md12 class="text-xs-center">
        <h1>Сводна информация </h1>
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex md12 class="mt-5">
        <line-example ></line-example>
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
      return {}
    },
    computed: {
      // loading() {
      //   return this.$store.getters.loading
      // },
      // getTemps() {
      //   return this.$store.getters.getTempJson[0]
      // }
    },
    //  created() {
    //    this.$store.dispatch('getTemp')
    // },


  }
</script>

<style>
  .small {
    max-width: 600px;
    margin: 150px auto;
  }

</style>
