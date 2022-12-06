<template>
  <div>
    <div class="wrapper">
      <!-- Sidebar Holder -->
      <nav id="sidebar">
        <div class="sidebar-header">
          <span class="text-decoration-none">
            <h1 class="mb-4 display-5 font-weight-semi-bold">
              <span class="text-secondary font-weight-bold border border-white px-3 mr-1">E</span>Knailk
            </h1>
          </span>
        </div>
        <ul class="list-unstyled components">
          <li>
            <NuxtLink to="/dashboard/home">
              <i class="fa fa-home"></i>
              Home
            </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/dashboard/user">
              <i class="fa fa-briefcase"></i>
              User
            </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/dashboard/product">
              <i class="fa fa-link"></i>
              Product
            </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/dashboard/add">
              <i class="fa fa-paperclip"></i>
              Add new product
            </NuxtLink>
          </li>
          <li>
            <a href="#pageSubmenu" data-toggle="collapse" aria-expanded="false">
              <i class="fa fa-file"></i>
              Pages
            </a>
            <ul class="collapse list-unstyled" id="pageSubmenu">
              <li><a href="#">Page 1</a></li>
              <li><a href="#">Page 2</a></li>
              <li><a href="#">Page 3</a></li>
            </ul>
          </li>
        </ul>
      </nav>

      <!-- Page Content Holder -->
      <div id="content">
        <nuxt-child />
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style scoped>
@import "~/assets/css/dashboard.css"
</style>
