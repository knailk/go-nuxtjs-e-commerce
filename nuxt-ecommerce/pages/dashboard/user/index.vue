<template>
  <div>
    <nav class="navbar navbar-default">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" id="sidebarCollapse" class="btn btn-info navbar-btn">
            <i class="fas fa-sign-out-alt" aria-hidden="true"></i>
            <span>Log Out</span>
          </button>
        </div>
      </div>
    </nav>
    <div class="container">
      <table class="table table-striped table-hover .table-responsive w-100 d-block d-md-table">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">ID</th>
            <th scope="col">Email</th>
            <th scope="col">Full Name</th>
            <th scope="col">Gender</th>
            <th scope="col">Phone</th>
            <th scope="col">Created</th>
            <th scope="col">Updated</th>
            <th scope="col">Status</th>
          </tr>
        </thead>
        <tbody class="align-middle">
          <tr v-for="(item, index) in userList" :key="item.id">
            <th scope="row">{{ (index + 1) }}</th>
            <td>{{ item.id }}</td>
            <td>{{ item.email }}</td>
            <td style="overflow: hidden;">{{ item.name }}</td>
            <td>{{ item.gender }}</td>
            <td>{{ item.phone }}</td>
            <td>{{ item.createdAt.split('T')[0] }}</td>
            <td>{{ item.updatedAt.split('T')[0] }}</td>
            <td><button type="button" class="btn"
                v-bind:class="{ 'bg-success': !item.isDeleted, 'bg-danger': item.isDeleted }"
                @click="changeStatus(item.id)"><i class="fa fa-times" aria-hidden="true"></i></button></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      filter: "email"
    }
  },
  async asyncData({ $axios }) {
    const userList = await $axios.$get("/admin/user", { params: { filter: "email" } });
    return { userList };
  },
  methods: {
    async changeStatus(id) {
      if (id != 2370508662) {
        await this.$axios.delete('/admin/user/' + id)
        this.$nuxt.refresh()
      }
    }
  },
  scrollToTop: true,
}
</script>

<style>

</style>
