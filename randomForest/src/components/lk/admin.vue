<template>
  <v-container>
    <v-layout>
      <v-flex md12 class="text-xs-center mb-5">
        <h1>Личный кабинет страховой компании</h1>
      </v-flex>
    </v-layout>
    <v-layout row>
      <v-flex md4>
        <v-card>
          <v-img
            src="https://cdn.vuetifyjs.com/images/lists/ali.png"
            height="300px"
            v-model="drawer"
          >
            <v-layout
              column
              fill-height
            >
              <v-card-title>
                <!--                <v-btn dark icon>-->
                <!--                  <v-icon>chevron_left</v-icon>-->
                <!--                </v-btn>-->

                <v-spacer></v-spacer>
                <!--              <app-update-modal :id="getUser._id"-->
                <!--                                :firstName="getUser.FirstName"-->
                <!--                                :secondName="getUser.SecondName"-->
                <!--                                :lastName="getUser.lastName"-->
                <!--                                :phone="getUser.mobile"-->
                <!--                                :email="getUser.email"-->
                <!--                                :workEmail="getUser.work_email"-->
                <!--                                :workPhone="getUser.phone"-->
                <!--                                :addres="getUser.address"-->
                <!--              ></app-update-modal>-->
              </v-card-title>
              <v-spacer></v-spacer>

              <v-card-title class="white--text pl-5 pt-5">
                              <div class="display-1 pl-5 pt-5">{{getCompanyOne.Name}}</div>
              </v-card-title>
            </v-layout>
          </v-img>
          <v-list two-line>

<!--            <v-list-tile @click="">-->
<!--              <v-list-tile-action>-->
<!--                <v-icon color="indigo"></v-icon>-->
<!--              </v-list-tile-action>-->

<!--              <v-list-tile-content>-->
<!--                <v-list-tile-title>{{getCompanyOne.Description}}</v-list-tile-title>-->
<!--                <v-list-tile-sub-title>Описание</v-list-tile-sub-title>-->
<!--              </v-list-tile-content>-->
<!--            </v-list-tile>-->
<!--            <v-divider inset></v-divider>-->

            <v-list-tile @click="">
              <v-list-tile-action>
                <v-icon color="indigo">phone</v-icon>
              </v-list-tile-action>

              <v-list-tile-content>
                <v-list-tile-title>{{getCompanyOne.Phone}}</v-list-tile-title>
                <v-list-tile-sub-title>Мобильный</v-list-tile-sub-title>
              </v-list-tile-content>

            </v-list-tile>

            <v-divider inset></v-divider>


            <v-list-tile @click="">
              <v-list-tile-action>
                <v-icon color="indigo">email</v-icon>
              </v-list-tile-action>

              <v-list-tile-content>
                <v-list-tile-title>{{getCompanyOne.Email}}</v-list-tile-title>
                <v-list-tile-sub-title>Рабочий email</v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>

            <v-divider inset></v-divider>

            <v-list-tile @click="">
              <v-list-tile-action>
                <v-icon color="indigo">location_on</v-icon>
              </v-list-tile-action>

              <v-list-tile-content>
                <v-list-tile-title>{{getCompanyOne.Address}}</v-list-tile-title>
                <v-list-tile-sub-title>Адрес</v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list>
        </v-card>
      </v-flex>

      <v-flex md8 style="margin-left: 50px">
        {{getCompanyOne.Description}}

      </v-flex>

    </v-layout>
    <v-tabs
      centered
      color="primary"
      style="margin-top: 50px"
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
              <v-btn small color="green" to="/client" style="color: white;">Посмотреть</v-btn>
            </td>
          </template>
        </v-data-table>
      </v-tab-item>

      <v-tab-item
        value="tab-2">
        <v-data-table
          :headers="headers2"
          :items="getNonClientAllJsons"
          class="elevation-1"
        >
          <template v-slot:items="props">
            <td>{{ props.item.ID }}</td>
            <td class="text-xs-center">{{ props.item.Name }}</td>
            <td class="text-xs-center">{{ props.item.DateEnd }}</td>
            <td class="text-xs-center">{{ props.item.Type }}</td>
            <td class="text-xs-center">{{ props.item.Contacts }}</td>
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
        headers2: [
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
          {text: 'Ознакомиться с компанией', value: '', align: 'center', sortable: false}

        ],
      }
    },
    computed: {
      getClientAllJsons() {
        if (this.$store.getters.getClientAllJson) {
          return this.$store.getters.getClientAllJson
        }
      },
      getNonClientAllJsons() {
        if (this.$store.getters.getNonClientAllJson) {
          return this.$store.getters.getNonClientAllJson
        }
      },
      getCompanyOne() {
        if (this.$store.getters.companyOnesJson) {
          console.log('sdsadsada')
          console.log(this.$store.getters.companyOnesJson)
          return this.$store.getters.companyOnesJson
        }
      },
    },
    created() {
      this.$store.dispatch('getClientAll')
      this.$store.dispatch('getComponyOneJson')

    }
  }
</script>

<style scoped>

</style>
