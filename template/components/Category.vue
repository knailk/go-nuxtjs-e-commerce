<template>
  <!-- Categories Start -->
  <div class="container-fluid pt-5">
        <div class="row px-xl-5 pb-3">
          <!--  -->
            <div class="col-lg-3 col-md-6 pb-1" v-for="item in catelist" :key="item.id">
                <div class="cat-item d-flex flex-column border mb-4" style="padding: 30px;">
                    <p class="text-right">{{item.numberProduct}} Products</p>
                    <!-- <NuxtLink :to="'/categories/'+ item.name.toLowerCase().replace(/\s/g, '-')" style="height: 80%" class="cat-img position-relative overflow-hidden mb-3"> -->
                      <NuxtLink :to="'/categories/'+ item.id" style="height: 80%" class="cat-img position-relative overflow-hidden mb-3">
                        <img style="max-width: 70%" class="img-fluid" :src="require(`~/assets/img/logo-`+ item.name.toLowerCase() +`.png`)" alt="">
                      </NuxtLink >
                    <h5 class="font-weight-semi-bold m-0">{{item.name}}</h5>
                </div>
            </div>
          <!--  -->
        </div>
    </div>
    <!-- Categories End -->
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const catelist = await $axios.$get('http://localhost:8081/product')
    return { catelist }
  },
  props: ['catelist' ],
}
</script>

<style>

</style>
