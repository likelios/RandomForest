<template>
  <v-container>
    <v-layout>
      <v-flex md12 class="text-xs-center">
        <h1>Сводна информация </h1>
      </v-flex>
    </v-layout>
    <v-layout>
      <line-example></line-example>
    </v-layout>

  </v-container>
</template>

<script>
  export default {
    name: "client",

    data() {
      return {
        data: [
          ["January", 10000],
          ["February", 12000],
          ["March", 18000],
          ["April", 11000],
          ["May", 9000]
        ]
      }
    },
    computed: {

    }

  }
</script>

<style scoped>
  .small {
    max-width: 600px;
    margin: 150px auto;
  }
  .Chart {
    background: #212733;
    border-radius: 15px;
    box-shadow: 0px 2px 15px rgba(25, 25, 25, 0.27);
    margin:  25px 0;
  }

  .Chart h2 {
    margin-top: 0;
    padding: 15px 0;
    color:  rgba(255, 0,0, 0.5);
    border-bottom: 1px solid #323d54;
  }
</style>
