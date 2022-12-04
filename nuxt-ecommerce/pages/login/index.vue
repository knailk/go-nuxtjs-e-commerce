<template>
  <div>
    <section class="img js-fullheight">
      <section class="ftco-section">
        <div class="container">
          <div class="row justify-content-center">
            <div class="col-md-6 text-center mb-5">
              <h2 class="heading-section">Login</h2>
            </div>
          </div>
          <div class="row justify-content-center">
            <div class="col-md-6 col-lg-4">
              <div class="login-wrap p-0">
                <h3 class="mb-4 text-center">Have an account?</h3>
                <form action="#" class="signin-form" @submit.prevent="userLogin">
                  <div class="form-group">
                    <input type="text" v-model="form.email" class="form-control" placeholder="Email address" name="email">
                  </div>
                  <div class="form-group">
                    <input v-if="showPassword" v-model="form.password" type="text" id="password-field"
                      class="form-control" placeholder="Password" required>
                    <input v-else type="password" v-model="form.password" id="password-field" class="form-control"
                      placeholder="Password" required>
                    <span toggle="#password-field" class="fa fa-fw field-icon toggle-password "
                      :class="{ 'fa-eye-slash': showPassword, 'fa-eye': !showPassword }" v-on:click="toggleShow"></span>
                  </div>
                  <div class="form-group">
                    <button type="submit" class="form-control btn btn-primary submit px-3">Sign In</button>
                  </div>
                  <div class="form-group d-md-flex">
                    <div class="w-50">
                      <NuxtLink to="/register" style="color: #fcceb6">Sign Up Here!</NuxtLink>
                    </div>
                    <div class="w-50 text-md-right">
                      <a href="#" style="color: #fcceb6">Forgot Password</a>
                    </div>
                  </div>
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
                </div><br />
              </div>
            </div>
          </div>
        </div>
      </section>
    </section>
    <notifications group="foo" width=400 height=700 />
  </div>
</template>

<script>
export default {
  data() {
    return {
      showPassword: false,
      password: null,
      form: {
        email: "a@bcd",
        password: "a@bcda@bcd",
      },
    };
  },
  computed: {
    buttonLabel() {
      return (this.showPassword) ? "Hide" : "Show";
    }
  },
  methods: {
    toggleShow() {
      this.showPassword = !this.showPassword;
    },
    async userLogin() {
      try {
        let response = await this.$auth.loginWith("local", {
          data: this.form,
        })
      } catch (err) {
        console.log(err)
        var stt = ""
        if (err.response.status === 204) {
          stt = 'Please fill all fields!'
        } else if (err.response.status === 404) {
          stt = 'Username does not exist!'
        } else if (err.response.status === 401) {
          stt = 'Incorrect password!'
        } else {
          stt = 'Login failed!'
        }
        this.$notify({
          group: 'foo',
          type: 'error',
          title: 'Error',
          text: stt,
        })
      }
      //this.$nuxt.refresh()
    }
  }
}
</script>

<style scoped>
@import "~/assets/css/login.css";
</style>
