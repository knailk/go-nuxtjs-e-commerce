<template>
  <section class="img js-fullheight">
    <section class="ftco-section">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-md-6 text-center mb-5">
            <h2 class="heading-section">Register</h2>
          </div>
        </div>
        <div class="row justify-content-center">
          <div class="col-md-6 col-lg-4">
            <div class="login-wrap p-0">
              <h3 class="mb-4 text-center">Get's started</h3>
              <form action="#" class="signin-form" @submit.prevent="userRegister">
                <div class="form-group">
                  <input type="text" v-model="form.fullName" class="form-control" placeholder="Full name" required>
                </div><!-- form-group// -->
                <div class="form-group">
                  <input type="email" v-model="form.email" class="form-control" placeholder="Email address" required>
                </div><!-- form-group// -->
                <div class="form-group d-md-flex">
                  <select class="form-control" style="max-width: 120px;">
                    <option style="color: black;" selected="">+84</option>
                    <option style="color: black;" value="1">+972</option>
                    <option style="color: black;" value="2">+198</option>
                    <option style="color: black;" value="3">+512</option>
                  </select>
                  <input name="" v-model="form.phone" class="form-control" placeholder="Phone number" type="text">
                </div> <!-- form-group// -->
                <div class="form-group d-md-flex" style="margin-bottom:0">
                  <input id="male" type="radio" name="gender" class=" w-25 form-control  text-left" value="Male"
                    v-model="form.gender">
                  <label for="male" class="w-25">Male</label>
                  <input id="female" type="radio" name="gender" class=" w-25 form-control  text-left" value="Female"
                    v-model="form.gender">
                  <label for="female" class="w-25">Female</label>
                </div><!-- form-group// -->
                <div class="form-group">
                  <input type="password" v-model="form.password" id="password-field" class="form-control"
                    placeholder="Password" required>
                </div><!-- form-group// -->
                <div class="form-group">
                  <input type="password" id="password-field-repeat" class="form-control" autocomplete="password-field"
                    placeholder="Repeat Password" required>
                </div><!-- form-group// -->
                <div class="form-group">
                  <button type="submit" class="form-control btn btn-primary submit px-3">Sign Up</button>
                </div><!-- form-group// -->
                <div class="form-group d-md-flex">
                  <div class="w-75">
                    <p class="text-right"> If you have an account?</p>
                  </div>
                  <div class="w-25 text-md-right">
                    <NuxtLink class="text-right" to="/login" style="color: #fcceb6">Login here</NuxtLink>
                  </div>
                </div><!-- form-group// -->
              </form>
              <p class="w-100 text-center">&mdash; Or Sign In With &mdash;</p>
              <div class="col-lg-6 text-center text-lg-right">
                <div class="d-inline-flex align-items-center">
                  <a class="text-primary pr-4" href="fb.com">
                    <i class="fab fa-facebook-f" style="font-size:30px"></i>
                  </a>
                  <a class="text-warning px-4" href="twitter.com">
                    <i class="fab fa-twitter" style="font-size:30px"></i>
                  </a>
                  <a class="text-danger px-4" href="instagram.com">
                    <i class="fab fa-instagram" style="font-size:30px"></i>
                  </a>
                  <a class="text-info px-4" href="https://www.linkedin.com/">
                    <i class="fab fa-linkedin-in" style="font-size:30px"></i>
                  </a>
                  <a class="text-success pl-4" href="youtube.com">
                    <i class="fab fa-youtube" style="font-size:30px"></i>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </section>
</template>

<script>
export default {
  auth: 'guest',
  data() {
    return {
      showPassword: false,
      password: null,
      form: {
        email: "",
        password: "",
        fullName: "",
        phone: "",
        gender: ""
      },
    };
  },
  computed: {
    buttonLabel() {
      return (this.showPassword) ? "Hide" : "Show";
    }
  },
  methods: {
    async userRegister() {
      await this.$axios.post("signup", this.form)
      try {
        let response = await this.$auth.loginWith("local", {
          data: this.form,
        });
        console.log(response)
        this.$router.push('/')
      } catch (err) {
        console.log(err)
        //   var stt =""
        //   if (err.response.status === 204) {
        //     stt = 'Please fill all fields!'
        //   } else if (err.response.status === 404) {
        //     stt = 'Username does not exist!'
        //   } else if (err.response.status === 401) {
        //     stt = 'Incorrect password!'
        //   }else if (err.response.status === 500) {
        //     stt = 'Login failed!'
        //   }
        //   this.$notify({
        //     group: 'foo',
        //     type: 'error',
        //     title: 'Error',
        //     text: stt,
        //   })
      }
      this.$nuxt.refresh()
    }
  }
}
</script>

<style scoped>
@import "~/assets/css/login.css";
</style>
