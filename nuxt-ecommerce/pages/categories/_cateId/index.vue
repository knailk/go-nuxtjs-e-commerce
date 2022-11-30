<template>
  <div>
    <TopBar></TopBar>
    <NarBar :cateList="catelist"></NarBar>
    <!-- Page Header Start -->
    <div class="container-fluid bg-secondary mb-5">
      <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
        <!-- <h1 class="font-weight-semi-bold text-uppercase mb-3">{{ listProducts.category }}</h1> -->
        <h1 class="font-weight-semi-bold text-uppercase mb-3"> {{listProducts.category}}</h1>
        <div class="d-inline-flex">
          <p class="m-0"><a href="">Home</a></p>
          <p class="m-0 px-2">-</p>
          <p class="m-0">Our Shop</p>
        </div>
      </div>
    </div>
    <!-- Page Header End -->
    <FilterShop :listProducts="listProducts"></FilterShop>
    <Footer></Footer>
  </div>

</template>

<script>
// import "@/assets/css/style.css";
export default {

  async asyncData({ $axios, params }) {
    const catelist = await $axios.$get('http://localhost:8081/categories')
    //const listProducts = await $axios.$get('/product/' + this.$route.params.cateId).data
    const listProducts = await $axios.$get('http://localhost:8081/product/' + params.cateId)
    return { catelist, listProducts }
  },
  methods: {
    toUpper(str) {
      return str.replace(/(?:^|\s|-)\S/g, x => x.toUpperCase()).replaceAll('-', ' ');
    }
  },
  mounted(){
  },
  scrollToTop: true

}
</script>

<style>

</style>
