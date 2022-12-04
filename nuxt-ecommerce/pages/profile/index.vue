<template>
<div>
  <TopBar :numberProductInCart = "numberProductInCart"/>
  <Profile/>
  <Footer />
</div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const listProductsInCart = await $axios.$get("/cart");
    let numberProductInCart;
    if (listProductsInCart.totalPrice == 0) numberProductInCart = 0
    else numberProductInCart = listProductsInCart.listProductsInCart.length
    return { numberProductInCart };
  },
  scrollToTop: true,
}
</script>

<style>

</style>

