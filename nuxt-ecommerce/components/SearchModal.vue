<template>
  <div>
    <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document" >
        <div class="modal-content">
          <div class="modal-header">
            <div class="input-group">
              <input type="text" class="form-control" placeholder="Search for products" style="border: solid #D19C97"
                @keyup="searchTimeOut()" v-model="inputQuery">
              <div class="input-group-append">
                <span class="input-group-text bg-transparent text-primary">
                  <i class="fa fa-search"></i>
                </span>
              </div>
            </div>
          </div>
          <div class="modal-body">
            <div class="container-fluid" style="overflow:scroll;height: 510px;">
              <NuxtLink to="#" style="display:block" v-for="item in listProducsByQuery" :key="item.id">
                <div class="row search-item">
                  <div class="col-4"><img class="w-100" :src="require(`~/assets/img/${item.image}`)" alt=""></div>
                  <div class="col-4">
                    <p>{{ item.name }}</p>
                  </div>
                  <div class="col-4">
                    <p>{{ item.price }}$</p>
                  </div>
                </div>
              </NuxtLink>

            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
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
      inputQuery: "",
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
      if (this.inputQuery != "") {
        const listProducs = await this.$axios.$get("/product", { params: { query: this.inputQuery } });
        this.listProducsByQuery = listProducs
      } else {
        this.listProducsByQuery = ""
      }

      console.log(this.listProducsByQuery)
      return { listProducs };
    }
  }
}
</script>

<style scoped>
.search-item {
  max-height: 100px;
  text-align: center;
}

.search-item img {
  object-fit: contain;
  max-height: 100px;
  height: 100%;
}

.search-item img:hover {
  transform: scale(1.2);
}

.col-4 {
  margin: auto;
}

a {
  color: black;
  border-bottom: #D19C97 ridge 1px;
  border-width: 75%;
}

a:hover {
  transform-origin: bottom left;
  color: #D19C97;
  text-decoration: none;
}
.container-fluid::-webkit-scrollbar {
  display: none;
}
</style>
