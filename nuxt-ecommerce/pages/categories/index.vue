<template>
  <div>
    <TopBar :numberProductInCart = "numberProductInCart"/>
    <NarBar :cateList="catelist"/>
    <Featured />
    <Category :cateList="catelist"/>
    <Footer />
  </div>
</template>

<script>
// import "@/assets/css/style.css";
export default {
  scrollToTop: true,
  async asyncData({ $axios }) {
    const catelist = await $axios.$get('/categories');
    const listProductsInCart = await $axios.$get("/cart");
    let numberProductInCart;
    if (listProductsInCart.totalPrice == 0) numberProductInCart = 0
    else numberProductInCart = listProductsInCart.listProductsInCart.length
    return { catelist, listProductsInCart, numberProductInCart }
  },
}
</script>

<style>

</style>
