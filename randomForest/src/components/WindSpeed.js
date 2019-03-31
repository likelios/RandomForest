import {Line} from 'vue-chartjs'

export default {
  extends: Line,
  props: {
    chartData: {
      type: Array ,
      required: false
    },
    chartLabels: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      datacollection: {
        //Data to be represented on x-axis
        labels:  this.chartLabels,
        datasets: [{
          label: 'Сила ветра',
          backgroundColor: '#f87979',
          pointBackgroundColor: 'white',
          borderWidth: 1,
          pointBorderColor: '#249EBF',
          //Data to be represented on y-axis
          data:  this.chartData,
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
    // getWindJsons() {
    //   this.$store.getters.getWindJson
    // }
  },
  mounted() {

    //renderChart function renders the chart with the datacollection and options object.
    this.renderChart(this.datacollection, this.options)
  },
  // created() {
  //   this.$store.dispatch('getWind')
  // }
}
