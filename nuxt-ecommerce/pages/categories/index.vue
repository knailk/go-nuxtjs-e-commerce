<template>
  <div>
    <TopBar />
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
    const catelist = await $axios.$get('http://localhost:8081/product')
    return { catelist }
  },
}
</script>

<style>

</style>
