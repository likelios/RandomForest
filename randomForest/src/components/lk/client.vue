<template>
  <v-container grid-list-md text-xs-center>

    <v-layout v-if="loading">
      <v-flex md12 class="text-xs-center pt-5">
        <v-progress-circular indeterminate :size="100" :width="4" color="purple">
        </v-progress-circular>
      </v-flex>
    </v-layout>

    <v-layout v-else>
      <v-flex md12 class="text-xs-center">
        <h1>Сводная информация </h1>
      </v-flex>
    </v-layout>

    <v-layout justify-space-between fill-height>
      <v-flex md6 class="mt-5">
        <line-example v-if="loaded" :chart-data="value" :chart-labels="label"></line-example>
        <v-divider class="mt-3 mb-3"></v-divider>
        <wind-example v-if="loaded" :chart-data="valueHumidity" :chart-labels="labelHumidity"></wind-example>
        <v-divider class="mt-3 mb-3"></v-divider>
        <humidity-example v-if="loaded" :chart-data="valueWind" :chart-labels="labelWind"></humidity-example>
      </v-flex>
      <v-flex md6 class="mt-5 ml-5">
        <h1>Информация для клиента</h1>
        <v-data-table
          :headers="headers"
          :items="getEvents"
          class="elevation-1"
        >
          <template v-slot:items="props">
            <td>{{ props.item.deviceID }}</td>
            <td class="text-xs-center">{{ props.item.coord }}</td>
            <td class="text-xs-center">{{ props.item.temp }}</td>
            <td class="text-xs-center">{{ props.item.humidity }}</td>
            <td class="text-xs-center">{{ props.item.windSpeed }}</td>
            <td class="text-xs-center yellow"
                v-if="props.item.temp > 35 && props.item.humidity < 30 ">
              Повышается вероятность засухи
            </td>
            <td class="text-xs-center" v-else> Все норм</td>
          </template>
        </v-data-table>
      </v-flex>
    </v-layout>
    <v-layout md12>


      <yandex-map
        :coords="[54.62896654088406, 39.731893822753904]"
        zoom="10"
        style="min-width: 1200px; height: 600px;"
        :cluster-options="{
    1: {clusterDisableClickZoom: true}
  }"
        :behaviors="['ruler']"
        :controls="['trafficControl']"
        :placemarks="placemarks"
        map-type="hybrid"
      >
        <!--        @map-was-initialized="initHandler"-->

        <!--        <ymap-marker-->
        <!--          marker-id="1"-->
        <!--          marker-type="placemark"-->
        <!--          :coords="[54.7, 39.7]"-->
        <!--          hint-content="Hint content 1"-->
        <!--          :balloon="{header: 'header', body: 'body', footer: 'footer'}"-->
        <!--          :icon="{color: 'green', glyph: 'cinema'}"-->
        <!--          cluster-name="1"-->
        <!--        ></ymap-marker>-->

        <!--        <ymap-marker-->
        <!--          marker-id="2"-->
        <!--          marker-type="placemark"-->
        <!--          :coords="[54.6, 39.8]"-->
        <!--          hint-content="Hint content 1"-->
        <!--          :balloon="{header: 'header', body: 'body', footer: 'footer'}"-->
        <!--          :icon="{color: 'green', glyph: 'cinema'}"-->
        <!--          cluster-name="1"-->
        <!--        ></ymap-marker>-->

        <!--      <ymap-marker-->
        <!--        marker-id="3"-->
        <!--        marker-type="circle"-->
        <!--        :coords="[54.62896654088406, 39.731893822753904]"-->
        <!--        circle-radius="1600"-->
        <!--        hint-content="Hint content 1"-->
        <!--        :marker-fill="{color: '#000000', opacity: 0.4}"-->
        <!--        :marker-stroke="{color: '#ff0000', width: 5}"-->
        <!--        :balloon="{header: 'header', body: 'body', footer: 'footer'}"-->
        <!--      ></ymap-marker>-->

      </yandex-map>
    </v-layout>
  </v-container>
</template>

<script>
  export default {
    name: "client",

    data() {
      return {
        placemarks: [
          {
            coords: [54.8, 39.8],
            properties: {}, // define properties here
            options: {}, // define options here
            clusterName: "1",
            balloonTemplate: '<div>"Тестовая карта"</div>',
            callbacks: {
              click: function () {
              }
            }
          }
        ],
        loaded: false,
        label: [],
        value: [],
        labelHumidity: [],
        valueHumidity: [],
        labelWind: [],
        valueWind: [],
        headers: [
          {
            text: '№ датчика',
            align: 'left',
            value: 'deviceID'
          },
          {text: 'Координаты', value: 'coord'},
          {text: 'Температура', value: 'temp'},
          {text: 'Влажность', value: 'humidity'},
          {text: 'Сила ветра', value: 'windSpeed'},
          {text: 'Вероятность засухи', value: '', sortable: false},
        ],
        desserts: this.getEvents,
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
      },
      getEvents() {
        let a = this.$store.getters.getEventJson;
        if (a) {
          return a
        }
      },
      async getJson() {
        this.loaded = false
        await this.$store.dispatch('getTemp').then(() => {
          this.getTemps
        })
        await this.$store.dispatch('getWind').then(() => {
          this.getWinds
        })
        await this.$store.dispatch('getHumidity').then(() => {
          this.getHumiditys
          // this.loaded = true
        })
        this.$store.dispatch('getEvent')
        this.loaded = true
      }
    },
    created() {
      this.getJson
      // this.loaded = false
      // this.$store.dispatch('getTemp').then(() => {
      //   this.getTemps
      // })
      // this.$store.dispatch('getWind').then(() => {
      //   this.getWinds
      // })
      // this.$store.dispatch('getHumidity').then(() => {
      //   this.getHumiditys
      //   // this.loaded = true
      // })
      // this.$store.dispatch('getEvent').then(() => {
      //   this.loaded = true
      // })
    },


  }
</script>

<style>
  .small {
    max-width: 600px;
    margin: 150px auto;
  }

</style>
