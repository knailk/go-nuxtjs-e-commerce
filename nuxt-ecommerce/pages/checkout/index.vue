<template>
  <div>
    <TopBar :numberProductInCart = "numberProductInCart"/>
    <NarBar :cateList="catelist"/>
    <!-- Page Header Start -->
    <div class="container-fluid bg-secondary mb-5">
      <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
        <h1 class="font-weight-semi-bold text-uppercase mb-3">CHECKOUT</h1>
        <div class="d-inline-flex">
          <p class="m-0"><a href="">Home</a></p>
          <p class="m-0 px-2">-</p>
          <p class="m-0">Checkout</p>
        </div>
      </div>
    </div>
    <!-- Page Header End -->
    <Checkout :listProductsInCart ="listProductsInCart"/>
    <Footer />
  </div>
</template>

<script>
// import "@/assets/css/style.css" ;
export default {
  async asyncData({ $axios }) {
    const catelist = await $axios.$get('http://localhost:8081/categories')
    const listProductsInCart = await $axios.$get("/cart");
    let numberProductInCart;
    if (listProductsInCart.totalPrice == 0) numberProductInCart = 0
    else numberProductInCart = listProductsInCart.listProductsInCart.length
    return { catelist, listProductsInCart, numberProductInCart }
  },
    scrollToTop: true
}
</script>

<style>

</style>
