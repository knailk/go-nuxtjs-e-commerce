<template>
  <!-- Page Header Start -->
  <div class="container-fluid bg-secondary mb-5">
    <div class="d-flex flex-column align-items-center justify-content-center" style="min-height: 300px">
      <h1 class="font-weight-semi-bold text-uppercase mb-3">Our Shop</h1>
      <div class="d-inline-flex">
        <p class="m-0"><NuxtLink to="/">Home</NuxtLink></p>
        <p class="m-0 px-2">-</p>
        <p class="m-0">{{cateName}}</p>
      </div>
    </div>
  </div>
  <!-- Page Header End -->
</template>

<script>
export default {
  mounted (){
    data = this.$route.params.cateName;
    cateName = data.replace(/(?:^|\s|-)\S/g, x => x.toUpperCase());
  },
  props: ['cateName']
};
</script>

<style>

</style>
