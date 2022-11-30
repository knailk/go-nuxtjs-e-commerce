<template>
  <div>
    <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <div class="input-group">
              <input type="text" class="form-control" placeholder="Search for products" style="border: solid #D19C97"
                @keyup="searchTimeOut()" v-model="inputQuery" >
              <div class="input-group-append">
                <span class="input-group-text bg-transparent text-primary">
                  <i class="fa fa-search"></i>
                </span>
              </div>
            </div>
          </div>
          <div class="modal-body">
            <div class="container-fluid" >
              <div class="row" v-for="item in listProducsByQuery" :key="item.id">
                <div class="col-lg-4">Toan an cuc</div>
                <div class="col-lg-4">{{item.name}}</div>
                <div class="col-lg-4">{{item.price}}</div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            <button type="button" class="btn btn-primary">Save changes</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      inputQuery:"",
      listProducsByQuery: ''
    }
  },
  methods: {
    searchTimeOut() {
      if (this.timer) {
        clearTimeout(this.timer);
        this.timer = null;
      }
      this.timer = setTimeout(() => {
        this.getProductsSearch()
      }, 800);
    },
    async getProductsSearch() {
      console.log(this.inputQuery)
      const listProducs = await this.$axios.$get("/product",{ params: { query: this.inputQuery }});
      this.listProducsByQuery = listProducs
      console.log(this.listProducsByQuery)
      return { listProducs };
    }
  }
}
</script>

<style scoped>

</style>
