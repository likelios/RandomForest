<template>
  <v-container text-xs-center>
    <v-layout>
      <v-flex md12>
        <h1>Список компаний</h1>
        <v-data-table
          :headers="headers"
          :items="getCompanyJson"
          class="elevation-1"
          prev-icon="mdi-menu-left"
          next-icon="mdi-menu-right"
          sort-icon="mdi-menu-down"
        >
          <template v-slot:items="props">
            <td>{{ props.item.id }}</td>
            <td class="text-xs-center">{{ props.item.name }}</td>
            <td class="text-xs-center">{{ props.item.rating }}</td>
          </template>
        </v-data-table>
      </v-flex>

    </v-layout>
  </v-container>
</template>

<script>
  export default {
    name: "Home",
    data() {
      return {
        headers: [
          {
            text: 'ID',
            align: 'center',
            sortable: false,
            value: 'id'
          },
          {text: 'Название', value: 'name', align: 'center', sortable: false,},
          {text: 'Рейтинг', value: 'rating', align: 'center'},
        ],
        // desserts:
        //   [
        //   {
        //     id: 1,
        //     name: 'Frozen Yogurt',
        //     rating: 159,
        //   },
        // ]
      }
    },
    computed: {
      loading() {
        return this.$store.getters.loading
      },
      getCompanyJson() {
        return this.$store.getters.companyJson
      },
    },
    created() {
      this.$store.dispatch('getComponyJson');

    }

  }
</script>

<style scoped>

</style>
