<template>
  <div>
    <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <div class="input-group">
              <input type="text" class="form-control" placeholder="Search for products" style="border: solid #D19C97"
                @keyup="searchTimeOut()">
              <div class="input-group-append">
                <span class="input-group-text bg-transparent text-primary">
                  <i class="fa fa-search"></i>
                </span>
              </div>
            </div>
          </div>
          <div class="modal-body">
            <div class="container-fluid">
              <div class="row">
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
              </div>
              <div class="row">
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
              </div>
              <div class="row">
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
              </div>
              <div class="row">
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
              </div>
              <div class="row">
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
                <div class="col-lg-4">.col-lg-4</div>
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
      const catelist = await this.$axios.$get("http://localhost:8081/product");
      return { catelist };
    }
  }
}
</script>

<style scoped>

</style>
