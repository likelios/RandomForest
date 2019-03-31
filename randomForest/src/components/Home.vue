<template>
  <v-container text-xs-center>
    <v-layout v-if="loading">
      <v-flex xs12 class="text-xs-center pt-5">
        <v-progress-circular indeterminate :size="100" :width="4" color="purple">
        </v-progress-circular>
      </v-flex>
    </v-layout>

    <v-layout v-else>
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
            <td>{{ props.item.ID }}</td>
            <td class="text-xs-center">{{ props.item.Name }}</td>
            <td class="text-xs-center">{{ props.item.Rating }}</td>
            <td class="text-xs-center">
              <v-btn small color="green" style="color: white;">Посмотреть</v-btn>
            </td>
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
            value: 'ID'
          },
          {text: 'Название', value: 'Name', align: 'center', sortable: false,},
          {text: 'Рейтинг', value: 'Rating', align: 'center'},
          {text: 'Ознакомиться с компанией', value: 'name', align: 'center', sortable: false}
        ],
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
