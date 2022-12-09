<template>
  <div>
    <notifications position="top right" width=400 height=700 group="foo" />
    <nav class="navbar navbar-default">
      <div class="container-fluid" style="justify-content:center">
        <div class="navbar-header" >
          <!-- <button type="button" id="sidebarCollapse" class="btn btn-info navbar-btn">
            <i class="fas fa-sign-out-alt" aria-hidden="true"></i>
            <span>Log Out</span>
          </button> -->
          <b >Create new product</b>
        </div>
      </div>
    </nav>
    <div class="container">
      <div class="row">
        <div class="col-md-4 mb-3">
          <div class="card">
            <div class="card-body">
              <div class="d-flex flex-column align-items-center text-center">
                <label for="img">Select image:</label>
                <input @change="onFileChange" type="file" id="img" name="img" accept="image/*">
                <img v-if="url" class="w-100 h-100" :src="url" alt="" width="150">
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
                  <input type="text" class="form-control" v-model="form.name" minlength="2" maxlength="50"
                    placeholder="Enter name of product">
                </div>
              </div>
              <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Description</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <textarea type="text" class="form-control" v-model="form.description" style="overflow-y:auto;"
                    placeholder="Product's description"></textarea>
                </div>
              </div>
              <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Available Units</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <input type="number" class="form-control" min=0 v-model="form.availableUnits" placeholder="Available">
                </div>
              </div>
              <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Quantity Sold</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <input type="number" class="form-control" min=0 v-model="form.quantitySold"
                    placeholder="Quantity sold">
                </div>
              </div>
              <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Price</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <input type="number" class="form-control" v-model="form.price" placeholder="Price">
                </div>
              </div>
              <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Category</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <select class="form-control" v-model="form.categoryId" placeholder="Select category">
                    <option v-for="item in cateList" :key="item.id" :value="item.id">{{ item.name }}</option>
                  </select>
                </div>
              </div>
              <!-- <div class="row mb-3">
                <div class="col-sm-3">
                  <h6 class="mb-0">Create From</h6>
                </div>
                <div class="col-sm-9 text-primary">
                  <input type="text" class="form-control" v-model="itemChoose.createdAt" readonly>
                </div>
              </div> -->
            </div>
          </div>
        </div>
      </div>
      <div class="row" style="justify-content:center">
        <button type="button" class="btn btn-info" style=" width:150px; margin-top: 20px;"
          @click="addNewProduct">Save</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware:  ['auth-admin'],
  data() {
    return {
      form: {
        name: null,
        price: null,
        description: null,
        quantitySold: null,
        availableUnits: null,
        image: null,
        categoryId: null,
      },
      url: null,
    }
  },
  async asyncData({ $axios }) {
    const cateList = await $axios.$get("/categories");
    return { cateList };
  },
  methods: {
    onFileChange(e) {
      if (e.target.files[0]) {
        this.form.image = e.target.files[0].name
        const file = e.target.files[0];
        this.url = URL.createObjectURL(file);
      }
    },
    async addNewProduct() {
      const isEmpty = Object.values(this.form).some(x => x === null || x === '');
      if (!isEmpty) {
        try {
          this.form.price = parseInt(this.form.price)
          this.form.quantitySold = parseInt(this.form.quantitySold)
          this.form.availableUnits = parseInt(this.form.availableUnits)
          await this.$axios.post("/product", this.form)
          this.$notify({
            group: 'foo',
            type: 'success',
            title: 'Create',
            text: "Create new product successfully!",
          })
        } catch (error) {
          this.$notify({
            group: 'foo',
            type: 'error',
            title: 'Error',
            text: error.toString(),
          })
        }
        this.$nuxt.refresh()
      }
      else{
        this.$notify({
            group: 'foo',
            type: 'error',
            title: 'Error',
            text: "Please fill out all field!",
          })
      }
    }
  }

}
</script>

<style>
.col-sm-3 {
  align-items: center;
  justify-content: left;
  display: flex
}
</style>
