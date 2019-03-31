<template>
  <v-container>
    <v-layout>
      <v-flex md12 class="text-xs-center">
        <h1>Личный кабинет страховой компании</h1>
      </v-flex>
    </v-layout>


    <v-tabs
      centered
      color="primary"
      dark
      icons-and-text
    >
      <v-tabs-slider color="yellow"></v-tabs-slider>

      <v-tab href="#tab-1">
        Клиенты
        <v-icon>favorite</v-icon>
      </v-tab>

      <v-tab href="#tab-2">
        Список хозяйств
        <v-icon>local_florist</v-icon>
      </v-tab>


      <v-tab-item
        key="1"
        value="tab-1"
      >
        <v-data-table
          :headers="headers"
          :items="getClientAllJsons"
          class="elevation-1"
        >
          <template v-slot:items="props">
            <td>{{ props.item.ID }}</td>
            <td class="text-xs-center">{{ props.item.Name }}</td>
            <td class="text-xs-center">{{ props.item.DateEnd }}</td>
            <td class="text-xs-center">{{ props.item.Type }}</td>
            <td class="text-xs-center">{{ props.item.Contacts }}</td>
            <td class="text-xs-center">{{ props.item.Sum }}</td>
            <td class="text-xs-center">
              <v-btn small color="green" style="color: white;">Посмотреть</v-btn>
            </td>
          </template>
        </v-data-table>
      </v-tab-item>

      <v-tab-item
        value="tab-2">
        <v-data-table
          :headers="headers"
          :items="desserts"
          class="elevation-1"
        >
          <template v-slot:items="props">
            <td>{{ props.item.ID }}</td>
            <td class="text-xs-center">{{ props.item.Name }}</td>
            <td class="text-xs-center">{{ props.item.DateEnd }}</td>
            <td class="text-xs-center">{{ props.item.Type }}</td>
            <td class="text-xs-center">{{ props.item.Contacts }}</td>
            <td class="text-xs-center">{{ props.item.Sum }}</td>
            <td class="text-xs-center">
              <v-btn small color="green" style="color: white;">Посмотреть</v-btn>
            </td>
          </template>
        </v-data-table>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>

<script>
  export default {
    name: "admin",
    data() {
      return {
        headers: [
          {
            text: 'ID',
            align: 'center',
            sortable: false,
            value: 'ID'
          },
          {text: 'Имя', value: 'Name', align: 'center', sortable: false},
          {text: 'Окончание полиса', value: 'DateEnd', align: 'center', sortable: false},
          {text: 'Юридическая форма', value: 'Type', align: 'center', sortable: false},
          {text: 'Контакты', value: 'Contacts', align: 'center', sortable: false},
          {text: 'Стоимость полиса', value: 'Sum', align: 'center'},
          {text: 'Ознакомиться с компанией', value: '', align: 'center', sortable: false}

        ],
        desserts: this.getClientAllJsons
      }
    },
    computed: {
     getClientAllJsons() {
        return this.$store.getters.getClientAllJson
      },
    },
    created() {
      this.$store.dispatch('getClientAll')
    }
  }
</script>

<style scoped>

</style>
