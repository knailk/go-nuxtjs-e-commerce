<template>

    <div>
      <TopBar></TopBar>
      <NarBar :catelist="catelist"></NarBar>
      <!-- Page Header Start -->
    <div class="container-fluid bg-secondary mb-5">
      <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
        <h1 class="font-weight-semi-bold text-uppercase mb-3">Product name</h1>
        <div class="d-inline-flex">
          <p class="m-0"><a href="">Home</a></p>
          <p class="m-0 px-2">-</p>
          <p class="m-0">Our Shop</p>
        </div>
      </div>
    </div>
    <!-- Page Header End -->
    <ProductDetail></ProductDetail>
    <AlsoLike></AlsoLike>
    <Footer></Footer>
    </div>
</template>

<script>
export default {
  async fetch() {
    this.catelist = await fetch(
        'http://localhost:8081/product'
      ).then(res => res.json()),
      this.productDetail = await fetch(
      'http://localhost:8081/product/' + this.$route.params.cateId +'/'+ this.$route.params.productId
    ).then(res => res.json())
  },
  scrollToTop: true
}
</script>

<style>

</style>
