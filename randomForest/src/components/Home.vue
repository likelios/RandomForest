<template>
 <h1>Главная</h1>
</template>

<script>
  export default {
    name: "Home",
    data() {
      return {

      }
    },
    computed:{
      loading() {
        return this.$store.getters.loading
      },
      categoryList(){
        return this.$store.getters.category
      }
    },
    created() {
      this.$store.dispatch('getCategoryJson')
    }
  }
</script>

<style scoped>

</style>
