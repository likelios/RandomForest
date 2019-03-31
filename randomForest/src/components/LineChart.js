import {Line} from 'vue-chartjs'

export default {
  extends: Line,
  props: ['getJson'],
  data() {
    return {
      datacollection: {
        //Data to be represented on x-axis
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
        datasets: [{
          label: 'Data One',
          backgroundColor: '#f87979',
          pointBackgroundColor: 'white',
          borderWidth: 1,
          pointBorderColor: '#249EBF',
          //Data to be represented on y-axis
          data: this.getJson
        }]
      },
      //Chart.js options that controls the appearance of the chart
      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true
            },
            gridLines: {
              display: true
            }
          }],
          xAxes: [{
            gridLines: {
              display: false
            }
          }]
        },
        legend: {
          display: true
        },
        responsive: true,
        maintainAspectRatio: false
      }
    }
  },
  computed: {
    // loading() {
    //   return this.$store.getters.loading
    // },
    getTemps() {
      console.log('13321')
      return this.$store.getters.getTempJson[0]
    }
  },
  mounted() {
    // console.log(this.getJson)
    //renderChart function renders the chart with the datacollection and options object.
    this.renderChart(this.datacollection, this.options)
  },
  created() {
    this.$store.dispatch('getTemp')
  },
}
