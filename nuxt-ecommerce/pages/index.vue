<template>
  <div>
    <TopBar />
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
    const catelist = await $axios.$get("http://localhost:8081/categories");
    const topProduct = await $axios.$get("http://localhost:8081/product/top");
    return { catelist, topProduct };
  },

  scrollToTop: true,
}
</script>
