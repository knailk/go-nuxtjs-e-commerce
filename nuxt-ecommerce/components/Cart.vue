<template>
  <div class="container-fluid pt-5">
    <div class="row px-xl-5">
      <div class="col-lg-8 table-responsive mb-5">
        <table class="table table-bordered text-center mb-0">
          <thead class="bg-secondary text-dark">
            <tr>
              <th>Products</th>
              <th>Price</th>
              <th>Quantity</th>
              <th>Total</th>
              <th>Remove</th>
            </tr>
          </thead>
          <tbody class="align-middle">
            <tr v-for="item in listProductsInCart.listProductsInCart" :key="item.productId">
              <td class="align-middle"><img :src="require(`~/assets/img/${item.image}`)" alt=""
                  style="width: 50px; height: 50px; text-align: center;"> {{ item.name }}</td>
              <td class="align-middle">${{ item.price }}</td>
              <td class="align-middle">
                <div class="input-group quantity mx-auto" style="width: 100px;">
                  <div class="input-group-btn">
                    <button class="btn btn-sm btn-primary btn-minus" @click="changeQuantity(item.productId, -1)">
                      <i class="fa fa-minus"></i>
                    </button>
                  </div>
                  <input type="text" class="form-control form-control-sm bg-secondary text-center"
                    :value="item.quantity">
                  <div class="input-group-btn">
                    <button class="btn btn-sm btn-primary btn-plus" @click="changeQuantity(item.productId, 1)">
                      <i class="fa fa-plus"></i>
                    </button>
                  </div>
                </div>
              </td>
              <td class="align-middle">${{ item.price * item.quantity }}</td>
              <td class="align-middle"><button class="btn btn-sm btn-primary"
                  @click="changeQuantity(item.productId, item.quantity)"><i class="fa fa-times"></i></button></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="col-lg-4">
        <form class="mb-5" action="">
          <div class="input-group">
            <input type="text" class="form-control p-4" placeholder="Coupon Code">
            <div class="input-group-append">
              <button class="btn btn-primary">Apply Coupon</button>
            </div>
          </div>
        </form>
        <div class="card border-secondary mb-5">
          <div class="card-header bg-secondary border-0">
            <h4 class="font-weight-semi-bold m-0">Cart Summary</h4>
          </div>
          <div class="card-body">
            <div class="d-flex justify-content-between mb-3 pt-1">
              <h6 class="font-weight-medium">Subtotal</h6>
              <h6 class="font-weight-medium">${{listProductsInCart.totalPrice}}</h6>
            </div>
            <div class="d-flex justify-content-between">
              <h6 class="font-weight-medium">Discount</h6>
              <h6 class="font-weight-medium">$0</h6>
            </div>
          </div>
          <div class="card-footer border-secondary bg-transparent">
            <div class="d-flex justify-content-between mt-2">
              <h5 class="font-weight-bold">Total</h5>
              <h5 class="font-weight-bold">${{listProductsInCart.totalPrice}}</h5>
            </div>
            <button class="btn btn-block btn-primary my-3 py-3">Proceed To Checkout</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['listProductsInCart'],
  data() {
    return {
      details: {
        productId: "",
        quantity: 0,
      },
    }
  },
  methods: {
    async changeQuantity(productId, quantity) {
      try {
        this.details.productId = productId.toString()
        this.details.quantity = quantity
        if (quantity == 1) {
          await this.$axios.$post('/cart/add', {
            productId: productId.toString(),
            quantity: 1,
          })
        } else {
          if (quantity == -1) this.details.quantity = 1
          await this.$axios.$post('/cart/remove', {
            productId: productId.toString(),
            quantity: quantity,
          })
        }
        this.$nuxt.refresh()
        await this.$auth.fetchUser()
      } catch (error) {
        console.log(error)
      }

    },
  },
}
</script>

<style>

</style>
