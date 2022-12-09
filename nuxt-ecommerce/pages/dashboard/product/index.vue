<template>
  <div>
    <div class="modal fade modify-product" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-xl">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">{{itemChoose.name}}</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-4 mb-3">
                <div class="card">
                  <div class="card-body">
                    <div class="d-flex flex-column align-items-center text-center">
                      <img class="w-100 h-100" :src="~/assets/img/itemChoose.image" alt="" width="150">
                    </div>
                  </div>
                </div>
              </div>
              <div class="col-md-8">
                <div class="card" style="display:flex">
                  <div class="card-body">
                    <div class="row mb-3">
                      <div class="col-sm-3">
                        <h6 class="mb-0">Name</h6>
                      </div>
                      <div class="col-sm-9 text-primary">
                        <input type="text" class="form-control" v-model="itemChoose.name" minlength="2" maxlength="50">
                      </div>
                    </div>
                    <div class="row mb-3">
                      <div class="col-sm-3">
                        <h6 class="mb-0">Description</h6>
                      </div>
                      <div class="col-sm-9 text-primary">
                        <textarea type="text" class="form-control" v-model="itemChoose.description" style="overflow-y:auto;"></textarea>
                      </div>
                    </div>
                    <div class="row mb-3">
                      <div class="col-sm-3">
                        <h6 class="mb-0">Available</h6>
                      </div>
                      <div class="col-sm-2 text-primary">
                        <input type="number" class="form-control" min=0 v-model="itemChoose.availableUnits">
                      </div>
                      <div class="col-sm-3">
                        <h6 class="mb-0">Sold</h6>
                      </div>
                      <div class="col-sm-4 text-primary">
                        <input type="number" class="form-control" min=0 v-model="itemChoose.quantitySold">
                      </div>
                    </div>
                    <div class="row mb-3">
                      <div class="col-sm-3">
                        <h6 class="mb-0">Price</h6>
                      </div>
                      <div class="col-sm-2 text-primary">
                        <input type="number" class="form-control" v-model="itemChoose.price" minlength="9" maxlength="11">
                      </div>
                      <div class="col-sm-3">
                        <h6 class="mb-0">Category</h6>
                      </div>
                      <div class="col-sm-4 text-primary">
                        <input type="text" class="form-control" v-model="currCategory" readonly>
                      </div>
                    </div>
                    <div class="row mb-3">
                      <div class="col-sm-3">
                        <h6 class="mb-0">Create From</h6>
                      </div>
                      <div class="col-sm-9 text-primary">
                        <input type="text" class="form-control" v-model="itemChoose.createdAt" readonly>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal" @click="deleteProduct(itemChoose.id)">Delete</button>
            <button type="button" class="btn btn-primary" data-dismiss="modal" @click="updateProduct()">Save</button>
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
              <button id="btn-img" data-toggle="modal" data-target=".modify-product" @click="sendItem(item)">
                <img class="img-fluid w-100" :src="require(`~/assets/img/${item.image}`)" alt="">
              </button>
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
  middleware:  ['auth-admin'],
  data() {
    return {
      currCategory: "Bag",
      productList: this.productList,
      itemChoose: {
        id:"",
        name:"",
        price:0,
        description:"",
        availableUnits:0,
        quantitySold:0,
        createdAt:""
      },
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
      this.itemChoose = item
      this.itemChoose.id = item.id.toString()
      console.log(this.item)
    },
    async updateProduct(){
      try {
        console.log(this.itemChoose)
        this.itemChoose.quantitySold = parseInt(this.itemChoose.quantitySold)
        this.itemChoose.availableUnits = parseInt(this.itemChoose.availableUnits)
        this.itemChoose.price = parseInt(this.itemChoose.price)
        const response = await this.$axios.post("/admin/product", this.itemChoose)
        this.$notify({
          group: 'foo',
          title: 'Notification',
          type: 'success',
          text: 'Update successful!',
          $: { enter: { opacity: [1, 0] }, leave: { opacity: [0, 1] } },
          ignoreDuplicates: true,
          width: 700
        })
      } catch (error) {
        this.$notify({
          group: 'foo',
          title: 'Notification',
          type: 'error',
          text: error.toString(),
          $: { enter: { opacity: [1, 0] }, leave: { opacity: [0, 1] } },
          ignoreDuplicates: true,
          width: 700
        })
      }
      this.$nuxt.refresh()
      await this.$auth.fetchUser()
    }
    // async deleteProduct(id) {
    //     await this.$axios.delete('/admin/product/' + id)
    //     this.$nuxt.refresh()
    //   }
  },
  scrollToTop: true,
}
</script>

<style lang="scss" scoped>

#btn-img{
  padding: 0;
  border: none;
  background: none;
}
#btn-img:focus { outline: none; }
.col-sm-3{
  align-items: center;
  justify-content: left;
  display:flex
}

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
