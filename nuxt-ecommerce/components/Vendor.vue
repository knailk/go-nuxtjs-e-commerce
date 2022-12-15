
<template>
  <!-- Vendor Start -->

  <div class="container-fluid py-5">
    <div class="row px-xl-5">
      <div class="col" v-for="i in 8" :key="i">
        <div class="vendor-carousel">
          <div class="vendor-item" >
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
.vendor-item img {
  width: 120px;
  border-style: none;
  background-image: linear-gradient(gray 100%, transparent 0);
}
</style>

