<template>
  <div>
    <notifications position="top right"  width=400 height=700 group="foo" />
    <!-- Also Like Products Start -->
    <div class="container-fluid py-5">
      <div class="text-center mb-4">
        <h2 class="section-title px-5"><span class="px-2">You May Also Like</span></h2>
      </div>
      <div class="row px-xl-5">
        <div class="row px-xl-5 pb-3">
          <div v-for="(item, index) in listProducts.products" :key="item.id" class="col-lg-3 col-md-6 col-sm-12 pb-1">
            <template v-if="index < 4">
              <div style="height: 100%" class="card product-item border-0">
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
                  <Button class="btn btn-sm text-dark p-0" @click="addToCart(item.id)"><i
                      class="fas fa-shopping-cart text-primary mr-1"></i>Add To Cart</Button>
                </div>
              </div>

            </template>
          </div>
        </div>
      </div>
    </div>
    <!-- Also Like Products End -->
  </div>
</template>

<script>
export default {
  props: ['listProducts'],
  methods: {
    async addToCart(id) {
      console.log(id.toString())
      await this.$axios.$post('/cart/add', {
        productId: id.toString(),
        quantity: 1,
      })
      await this.$auth.fetchUser()
      this.$notify({
        group: 'foo',
        title: 'Notification',
        type: 'success',
        text: 'Added to Cart!',
        $: {enter: {opacity: [1, 0]}, leave: {opacity: [0, 1]}},
        ignoreDuplicates: true,
        width:  700
      })
    }
  }
}
</script>

<style>

</style>
