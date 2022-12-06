<template>
  <div>
    <TopBar :numberProductInCart = "numberProductInCart"/>
    <NarBar :cateList="catelist" :isHome=true />
    <notifications group="foo" width=400 height=700 />
    <Featured />
    <Category :cateList="catelist" />
    <Offer />
    <TopProduct :topProduct="topProduct" />
    <Subscribe />
    <Vendor />
    <Footer />
  </div>
</template>

<script>
import "@/assets/css/style.css";
export default {
  async asyncData({ $axios }) {
    const catelist = await $axios.$get("/categories");
    const topProduct = await $axios.$get("/product/top");
    const listProductsInCart = await $axios.$get("/cart");
    let numberProductInCart;
    if (listProductsInCart.totalPrice == 0) numberProductInCart = 0
    else numberProductInCart = listProductsInCart.listProductsInCart.length
    return { catelist, topProduct, listProductsInCart, numberProductInCart };
  },
  scrollToTop: true,
}
</script>
