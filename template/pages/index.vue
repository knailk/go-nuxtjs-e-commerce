<template>
  <div>
    <TopBar />
    <NarBar :catelist="catelist"/>
    <Featured />
    <Category :catelist="catelist"/>
    <Offer />
    <TopProduct :topProduct="topProduct"/>
    <Subscribe />
     <Vendor />
    <Footer />
  </div>
</template>

<script>
export default {
  // data() {
  //     return {
  //       catelist: [],
  //       topProduct:[],
  //     }
  //   },
    async fetch() {
      this.catelist = await fetch(
        'http://localhost:8081/product'
      ).then(res => res.json()),
      this.topProduct = await fetch(
        'http://localhost:8081/product/top'
      ).then(res => res.json())
    },
    scrollToTop: true
}
</script>
