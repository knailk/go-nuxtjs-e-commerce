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
                <div class="form-group">
                  <input type="password" v-model="form.password" id="password-field" class="form-control"
                    placeholder="Password" required>
                </div><!-- form-group// -->
                <div class="form-group">
                  <input type="password" id="password-field-repeat" class="form-control" autocomplete="password-field" placeholder="Repeat Password"
                    required>
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
  auth: false,
  data() {
    return {
      showPassword: false,
      password: null,
      form: {
        email: "",
        password: "",
        fullName: "",
        phone: "",
        gender: "Male"
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
      try {
        await this.$axios.post("signup", this.form)
        let response = await this.$auth.loginWith("local", {
          data: this.form,
        });
        console.log(response)
        console.log(this.user)
        this.$axios.defaults.headers.common.Authorization = `${this.$auth.getToken(
          "local"
        )}`;
        this.$notify({
          group: 'foo',
          type: 'success',
          title: 'Authorization',
          text: "Register successful",
        })
      } catch (err) {
        if (err.response.status === 204) {
          this.status = 'Please fill all fields!'
        } else if (err.response.status === 404) {
          this.status = 'Username does not exist!'
        } else if (err.response.status === 401) {
          this.status = 'Incorrect password!'
        }else if (err.response.status === 500) {
          this.status = 'Login failed!'
        }
        this.$notify({
          group: 'foo',
          type: 'error',
          title: 'Error',
          text: this.status,
        })
      }
      //this.$nuxt.refresh()
    }
  }
}
</script>

<style scoped>
/* @import "~/static/css/login.css"; */
.gradient-custom {
  /* fallback for old browsers */
  background: #9A616D;

  background: -webkit-linear-gradient(to right, rgb(212, 216, 163), rgb(246, 186, 222));

  background: linear-gradient(to right, rgb(212, 216, 163), rgb(246, 186, 222))
}

.js-fullheight {
  background-image: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url('/bg.jpg');
  width: 100vw;
  height: 100vh;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}


/* ***************************** */
.justify-content-center {
  -webkit-box-pack: center !important;
  -ms-flex-pack: center !important;
  justify-content: center !important;
}

.row {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  -ms-flex-wrap: wrap;
  flex-wrap: wrap;
  margin-right: -15px;
  margin-left: -15px;
}

.col-md-6 {
  -webkit-box-flex: 0;
  -ms-flex: 0 0 50%;
  flex: 0 0 50%;
  max-width: 50%;
}

.col-lg-4 {
  -webkit-box-flex: 0;
  -ms-flex: 0 0 33.33333%;
  flex: 0 0 33.33333%;
  max-width: 33.33333%;
}

.login-wrap {
  position: relative;
  color: rgba(255, 255, 255, 0.9);
}

.login-wrap h3 {
  font-weight: 300;
  color: #fff;
}

.login-wrap .social {
  width: 100%;
}

.login-wrap .social a {
  width: 100%;
  display: block;
  border: 1px solid rgba(255, 255, 255, 0.4);
  color: #000;
  background: #fff;
}

.login-wrap .social a:hover {
  background: #000;
  color: #fff;
  border-color: #000;
}

.form-group {
  position: relative;
}

textarea.form-control {
  height: inherit !important;
}

.form-control {
  display: block;
  width: 100%;
  height: calc(1.5em + 0.75rem + 2px);
  padding: 0.375rem 0.75rem;
  font-size: 1rem;
  font-weight: 400;
  line-height: 1.5;
  color: #495057;
  background-color: #fff;
  background-clip: padding-box;
  border: 1px solid #ced4da;
  border-radius: 0.25rem;
  -webkit-transition: border-color 0.15s ease-in-out, -webkit-box-shadow 0.15s ease-in-out;
  transition: border-color 0.15s ease-in-out, -webkit-box-shadow 0.15s ease-in-out;
  -o-transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
  transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
  transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out, -webkit-box-shadow 0.15s ease-in-out;
}

@media (prefers-reduced-motion: reduce) {
  .form-control {
    -webkit-transition: none;
    -o-transition: none;
    transition: none;
  }
}

.form-control::-ms-expand {
  background-color: transparent;
  border: 0;
}

.form-control:focus {
  color: #495057;
  background-color: #fff;
  border-color: #80bdff;
  outline: 0;
  -webkit-box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

.form-control::-webkit-input-placeholder {
  color: #6c757d;
  opacity: 1;
}

.form-control:-ms-input-placeholder {
  color: #6c757d;
  opacity: 1;
}

.form-control::-ms-input-placeholder {
  color: #6c757d;
  opacity: 1;
}

.form-control::placeholder {
  color: #6c757d;
  opacity: 1;
}

.form-control:disabled,
.form-control[readonly] {
  background-color: #e9ecef;
  opacity: 1;
}

select.form-control:focus::-ms-value {
  color: #495057;
  background-color: #fff;
}

.form-control-file,
.form-control-range {
  display: block;
  width: 100%;
}

.form-control {
  background: transparent;
  border: none;
  height: 50px;
  color: white !important;
  border: 1px solid transparent;
  background: rgba(255, 255, 255, 0.08);
  border-radius: 40px;
  padding-left: 20px;
  padding-right: 20px;
  -webkit-transition: 0.3s;
  -o-transition: 0.3s;
  transition: 0.3s;
}

@media (prefers-reduced-motion: reduce) {
  .form-control {
    -webkit-transition: none;
    -o-transition: none;
    transition: none;
  }
}

.form-control::-webkit-input-placeholder {
  /* Chrome/Opera/Safari */
  color: rgba(255, 255, 255, 0.8) !important;
}

.form-control::-moz-placeholder {
  /* Firefox 19+ */
  color: rgba(255, 255, 255, 0.8) !important;
}

.form-control:-ms-input-placeholder {
  /* IE 10+ */
  color: rgba(255, 255, 255, 0.8) !important;
}

.form-control:-moz-placeholder {
  /* Firefox 18- */
  color: rgba(255, 255, 255, 0.8) !important;
}

.form-control:hover,
.form-control:focus {
  background: transparent;
  outline: none;
  -webkit-box-shadow: none;
  box-shadow: none;
  border-color: rgba(255, 255, 255, 0.4);
}

.form-control:focus {
  border-color: rgba(255, 255, 255, 0.4);
}

.d-md-flex {
  display: -webkit-box !important;
  display: -ms-flexbox !important;
  display: flex !important;
}

.was-validated .form-control:valid,
.form-control.is-valid {
  border-color: #28a745;
  padding-right: calc(1.5em + 0.75rem);
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 8 8'%3e%3cpath fill='%2328a745' d='M2.3 6.73L.6 4.53c-.4-1.04.46-1.4 1.1-.8l1.1 1.4 3.4-3.8c.6-.63 1.6-.27 1.2.7l-4 4.6c-.43.5-.8.4-1.1.1z'/%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: center right calc(0.375em + 0.1875rem);
  background-size: calc(0.75em + 0.375rem) calc(0.75em + 0.375rem);
}

.was-validated .form-control:valid:focus,
.form-control.is-valid:focus {
  border-color: #28a745;
  -webkit-box-shadow: 0 0 0 0.2rem rgba(40, 167, 69, 0.25);
  box-shadow: 0 0 0 0.2rem rgba(40, 167, 69, 0.25);
}

.was-validated .form-control:valid~.valid-feedback,
.was-validated .form-control:valid~.valid-tooltip,
.form-control.is-valid~.valid-feedback,
.form-control.is-valid~.valid-tooltip {
  display: block;
}

.was-validated textarea.form-control:valid,
textarea.form-control.is-valid {
  padding-right: calc(1.5em + 0.75rem);
  background-position: top calc(0.375em + 0.1875rem) right calc(0.375em + 0.1875rem);
}

.was-validated .form-control-file:valid~.valid-feedback,
.was-validated .form-control-file:valid~.valid-tooltip,
.form-control-file.is-valid~.valid-feedback,
.form-control-file.is-valid~.valid-tooltip {
  display: block;
}


.was-validated .form-control:invalid~.invalid-feedback,
.was-validated .form-control:invalid~.invalid-tooltip,
.form-control.is-invalid~.invalid-feedback,
.form-control.is-invalid~.invalid-tooltip {
  display: block;
}

.field-icon {
  position: absolute;
  top: 50%;
  right: 15px;
  -webkit-transform: translateY(-50%);
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
  color: rgba(255, 255, 255, 0.9);
}

.checkbox-wrap {
  display: block;
  position: relative;
  padding-left: 30px;
  margin-bottom: 12px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

input[type="checkbox"] {
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  padding: 0;
}

.btn-group-toggle>.btn input[type="checkbox"],
.btn-group-toggle>.btn-group>.btn input[type="radio"],
.btn-group-toggle>.btn-group>.btn input[type="checkbox"] {
  position: absolute;
  clip: rect(0, 0, 0, 0);
  pointer-events: none;
}

/* Hide the browser's default checkbox */
.checkbox-wrap input {
  position: absolute;
  opacity: 0;
  cursor: pointer;
  height: 0;
  width: 0;
}

/* Create a custom checkbox */
.checkmark {
  position: absolute;
  top: 0;
  left: 0;
}

/* Create the checkmark/indicator (hidden when not checked) */
.checkmark:after {
  content: "\f0c8";
  font-family: "FontAwesome";
  position: absolute;
  color: rgba(255, 255, 255, 0.1);
  font-size: 20px;
  margin-top: -4px;
  -webkit-transition: 0.3s;
  -o-transition: 0.3s;
  transition: 0.3s;
}

@media (prefers-reduced-motion: reduce) {
  .checkmark:after {
    -webkit-transition: none;
    -o-transition: none;
    transition: none;
  }
}

/* Show the checkmark when checked */
.checkbox-wrap input:checked~.checkmark:after {
  display: block;
  content: "\f14a";
  font-family: "FontAwesome";
  color: rgba(0, 0, 0, 0.2);
}

/* Style the checkmark/indicator */
.checkbox-primary {
  color: #fbceb5;
}

.checkbox-primary input:checked~.checkmark:after {
  color: #fbceb5;
}

.btn {
  cursor: pointer;
  border-radius: 40px;
  -webkit-box-shadow: none !important;
  box-shadow: none !important;
  font-size: 15px;
  text-transform: uppercase;
}

.btn:hover,
.btn:active,
.btn:focus {
  outline: none;
}

.btn.btn-primary {
  background: #fbceb5 !important;
  border: 1px solid #fbceb5 !important;
  color: #000 !important;
}

.btn.btn-primary:hover {
  border: 1px solid #fbceb5;
  background: transparent;
  color: #fbceb5;
}

.btn.btn-primary.btn-outline-primary {
  border: 1px solid #fbceb5;
  background: transparent;
  color: #fbceb5;
}

.btn.btn-primary.btn-outline-primary:hover {
  border: 1px solid transparent;
  background: #fbceb5;
  color: #fff;
}

@media (prefers-reduced-motion: reduce) {
  .checkmark:after {
    -webkit-transition: none;
    -o-transition: none;
    transition: none;
  }
}

.d-inline-flex {
  display: -webkit-inline-box !important;
  display: -ms-inline-flexbox !important;
  display: inline-flex !important;
}

.align-items-center {
  -webkit-box-align: center !important;
  -ms-flex-align: center !important;
  align-items: center !important;
}

.ftco-section {
  padding-top: 2em;
}

/* ****************************** */
</style>
