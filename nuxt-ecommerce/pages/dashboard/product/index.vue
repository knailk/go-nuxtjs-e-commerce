<template>
  <div>
    <div class="modal fade modify-product" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">{{item.name}}</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form>
              <div class="form-group">
                <label for="recipient-name" class="col-form-label">Recipient:</label>
                <input type="text" class="form-control" id="recipient-name">
              </div>
              <div class="form-group">
                <label for="message-text" class="col-form-label">Message:</label>
                <textarea class="form-control" id="message-text"></textarea>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            <button type="button" class="btn btn-primary">Send message</button>
          </div>
        </div>
      </div>
    </div>
    <nav class="navbar navbar-default">
      <div class="container-fluid">
        <div class="navbar-header">
          <div class="input-group-prepend">
            <button class="btn btn-outline-primary dropdown-toggle" type="button" data-toggle="dropdown"
              aria-haspopup="false" aria-expanded="true">{{ this.currCategory }}</button>
            <div class="dropdown-menu">
              <button class="dropdown-item" v-for="item in cateList" :key="item.id"
                @click="changeCategory(item.id, item.name)">{{ item.name }}</button>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div class="container" style="margin-top:20px">
      <div class="row pb-3">
        <div class="col-lg-3 col-md-6 col-sm-12 pb-1" v-for="item in productList" :key="item.id">
          <div style="height: 100%;" class="card product-item border-0 mb-4">
            <div div style="height: 15%" class="d-flex justify-content-end bg-light border">
              <button id="close" class="btn btn-lg">
                <i class="fa fa-times-circle"
                  v-bind:class="{ 'text-success': !item.isDeleted, 'text-danger': item.isDeleted }"></i>
              </button>
            </div>
            <div style="height: 75%"
              class="card-header product-img position-relative overflow-hidden bg-transparent border p-0">
              <button data-toggle="modal" data-target=".modify-product" @click="sendItem(item)"><img class="img-fluid w-100" :src="require(`~/assets/img/${item.image}`)" alt=""></button>
            </div>
            <div style="height: 20%" class="card-body border-left border-right text-center p-0">
              <h6 class="text-truncate">{{ item.name }}</h6>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currCategory: "Bag",
      productList: this.productList,
      item:""
    }
  },
  async asyncData({ $axios }) {
    const cateList = await $axios.$get("/categories");
    const productList = await $axios.$get("/admin/product/1");
    return { cateList, productList };
  },
  methods: {
    async changeCategory(id, name) {
      if (name != this.currCategory) {
        this.currCategory = name;
        this.productList = await this.$axios.$get("/admin/product/" + id);
      }
    },
    sendItem(item){
      this.item = item
    }
    // async changeStatus(id) {
    //     await this.$axios.delete('/admin/product/' + id)
    //     this.$nuxt.refresh()
    //   }
  },
  scrollToTop: true,
}
</script>

<style lang="scss" scoped>
.modify-product
#close {
  position: absolute;
  -webkit-transition: -webkit-transform .25s, opacity .25s;
  -moz-transition: -moz-transform .25s, opacity .25s;
  transition: transform .25s, opacity .25s;
  opacity: .75;
}

#close:hover {
  -webkit-transform: rotate(270deg);
  -moz-transform: rotate(270deg);
  transform: rotate(270deg);
  opacity: 1;
}
</style>
