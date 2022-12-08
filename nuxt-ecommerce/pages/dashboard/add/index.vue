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
      <div class="row">
        <div class="col-md-4 mb-3">
          <div class="card">
            <div class="card-body">
              <div class="d-flex flex-column align-items-center text-center">
                <img class="w-100 h-100" src="(~/assets/img/user.png)" alt="" width="150">
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
                  <textarea type="text" class="form-control" v-model="itemChoose.description"
                    style="overflow-y:auto;"></textarea>
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
  </div>
</template>

<script>
export default {
  data(){
    return {
      form:{
        
      }
    }
  }

}
</script>

<style>

</style>
