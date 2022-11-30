<template>

  <div>
    <TopBar></TopBar>
    <NarBar :cateList="catelist"></NarBar>
    <!-- Page Header Start -->
    <div class="container-fluid bg-secondary mb-5">
      <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
        <h1 class="font-weight-semi-bold text-uppercase mb-3">{{productDetail.name}}</h1>
        <div class="d-inline-flex">
          <p class="m-0"><a href="">Home</a></p>
          <p class="m-0 px-2">-</p>
          <p class="m-0">Our Shop</p>
        </div>
      </div>
    </div>
    <!-- Page Header End -->
    <ProductDetail :productDetail="productDetail" />
    <AlsoLike :listProducts="listProducts" />
    <Footer></Footer>
  </div>
</template>

<script>
// import "@/assets/css/style.css";
export default {
  scrollToTop: true,
  async asyncData({ $axios,params }) {
    const catelist = await $axios.$get('http://localhost:8081/product')
    const listProducts = await $axios.$get('http://localhost:8081/product/' + params.cateId)
    const productDetail = await $axios.$get('http://localhost:8081/product/' + params.cateId + '/' + params.productId)
    return { catelist, listProducts, productDetail }
  },
}
</script>

<style>

</style>
