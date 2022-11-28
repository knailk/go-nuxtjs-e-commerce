<template>
  <!-- Top Products Start -->
  <div class="container-fluid pt-5">
    <div class="text-center mb-4">
      <h2 class="section-title px-5"><span class="px-2">Top Products</span></h2>
    </div>
    <div class="row px-xl-5 pb-3">
      <div class="col-lg-3 col-md-6 col-sm-12 pb-1" v-for="item in topProduct" :key="item.id">
        <div style="height: 100%" class="card product-item border-0 mb-4">
          <div style="height: 70%"
            class="card-header product-img position-relative overflow-hidden bg-transparent border p-0">
            <img class="img-fluid w-100" :src="require(`~/assets/img/${item.image}`)" alt="">
          </div>
          <div style="height: 20%" class="card-body border-left border-right text-center p-0 pt-4 pb-3">
            <h6 class="text-truncate mb-3">{{ item.name }}</h6>
            <div class="d-flex justify-content-center">
              <h6>${{ item.price }}</h6>
              <h6 class="text-muted ml-2"><del>${{ item.price }}</del></h6>
            </div>
          </div>
          <div style="height: 10%" class="card-footer d-flex justify-content-between bg-light border">
            <NuxtLink :to="'/categories/' + item.category + '/' + item.id" class="btn btn-sm text-dark p-0"><i
                class="fas fa-eye text-primary mr-1"></i>View Detail</NuxtLink>
            <a href="" class="btn btn-sm text-dark p-0"><i class="fas fa-shopping-cart text-primary mr-1"></i>Add To
              Cart</a>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- Top Products End -->
</template>

<script>
export default {
  props:['topProduct'],
}
</script>

<style>

</style>
