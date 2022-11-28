<template>
  <div>
    <TopBar />
    <NarBar :cateList="catelist"/>
    <!-- Page Header Start -->
    <div class="container-fluid bg-secondary mb-5">
      <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
        <h1 class="font-weight-semi-bold text-uppercase mb-3">CONTACT US</h1>
        <div class="d-inline-flex">
          <p class="m-0"><a href="">Home</a></p>
          <p class="m-0 px-2">-</p>
          <p class="m-0">Contact</p>
        </div>
      </div>
    </div>
    <!-- Page Header End -->
    <Contact />
    <Footer />
  </div>
</template>

<script>
// import "@/assets/css/style.css" ;
export default {
  async asyncData({ $axios }) {
    const catelist = await $axios.$get('http://localhost:8081/product')
    return { catelist }
  },
    scrollToTop: true
}
</script>

<style>

</style>
