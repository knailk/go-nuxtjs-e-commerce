<template>
  <!-- Topbar Start -->
  <div class="container-fluid">
    <div class="row bg-secondary py-2 px-xl-5">
      <div class="col-lg-6 d-none d-lg-block">
        <div class="d-inline-flex align-items-center">
          <NuxtLink class="text-dark" to="#">FAQs</NuxtLink>
          <span class="text-muted px-2">|</span>
          <NuxtLink class="text-dark" to="#">Help</NuxtLink>
          <span class="text-muted px-2">|</span>
          <NuxtLink class="text-dark" to="#">Support</NuxtLink>
        </div>
      </div>
      <div class="col-lg-6 text-center text-lg-right">
        <div class="d-inline-flex align-items-center">
          <a class="text-dark px-2" href="fb.com">
            <i class="fab fa-facebook-f"></i>
          </a>
          <a class="text-dark px-2" href="twitter.com">
            <i class="fab fa-twitter"></i>
          </a>
          <a class="text-dark px-2" href="https://www.linkedin.com/">
            <i class="fab fa-linkedin-in"></i>
          </a>
          <a class="text-dark px-2" href="instagram.com">
            <i class="fab fa-instagram"></i>
          </a>
          <a class="text-dark pl-2" href="youtube.com">
            <i class="fab fa-youtube"></i>
          </a>
        </div>
      </div>
    </div>
    <div class="row align-items-center py-3 px-xl-5">
      <div class="col-lg-3 d-none d-lg-block">
        <NuxtLink to="/" class="text-decoration-none">
          <h1 class="m-0 display-5 font-weight-semi-bold"><span
              class="text-primary font-weight-bold border px-3 mr-1">E</span>Knailk</h1>
        </NuxtLink>
      </div>
      <div class="col-lg-6 col-6 text-left">
        <form action="">
          <div class="input-group">
            <input type="text" class="form-control" placeholder="Search for products" data-toggle="modal"
              data-target="#exampleModal">
            <div class="input-group-append">
              <span class="input-group-text bg-transparent text-primary">
                <i class="fa fa-search"></i>
              </span>
            </div>
          </div>
        </form>
      </div>
      <!-- Modal -->
      <SearchModal />
      <!-- Modal end -->
      <div class="col-lg-3 col-6 text-right">
        <a href="" class="btn border">
          <i class="fas fa-heart text-primary"></i>
          <span class="badge">9</span>
        </a>
        <NuxtLink to="/cart" class="btn border">
          <i class="fas fa-shopping-cart text-primary"></i>
          <span class="badge">0</span>
        </NuxtLink>
      </div>
    </div>
  </div>
  <!-- Topbar End -->
</template>

<script>
import { mapMutations } from 'vuex'
export default {
  methods: {
    ...mapMutations([
      'changeStateShowModal',
    ]),
  }
}
</script>

<style>

</style>
