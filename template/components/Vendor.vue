
<template>
  <!-- Vendor Start -->

  <div class="container-fluid py-5">
    <div class="row px-xl-5">
      <div class="col">
        <div class="owl-carousel vendor-carousel">
            <carousel
            :autoplay="true"
            :loop="true"
            :margin=29
            :nav="false"
            :autoplaySpeed=1000
            :responsive="{
              0: {
                items: 2,
              },
              576: {
                items: 3,
              },
              768: {
                items: 4,
              },
              992: {
                items: 5,
              },
              1200: {
                items: 6,
              }
            }"
          >
            <slide class="vendor-item border p-4" v-for="i in 8" :key="i" >
              <img :src="require(`~/assets/img/vendor-${i}.jpg`)" alt="" />
            </slide>
          </carousel>
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

