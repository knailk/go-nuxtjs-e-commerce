
<template>
  <!-- Vendor Start -->

  <div class="container-fluid py-5">
    <div class="row px-xl-5">
      <div class="col">
        <div class="owl-carousel vendor-carousel">
          <div class="vendor-item border p-4" v-for="i in 8" :key="i">
            <img :src="require(`~/assets/img/vendor-${i}.jpg`)" alt="" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Vendor End -->
</template>

<script>
export default {
};
</script>

<style lang="scss">
.vendor-carousel {
  height: 100px;

  .VueCarousel-slide {
    text-align: center;
  }
}

.vendor-item img {
  width: 120px;
  border-style: none;
  background-image: linear-gradient(gray 100%, transparent 0);
}
</style>

