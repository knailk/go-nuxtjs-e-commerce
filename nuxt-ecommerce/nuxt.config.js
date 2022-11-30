import webpack from "webpack";
export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'nuxt-ecommerce',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: "preconnect", href: "https://fonts.gstatic.com" },
      {
        rel: "stylesheet",
        href: "https://fonts.googleapis.com/css2?family=Poppins:wght@100;200;300;400;500;600;700;800;900&display=swap",
      },
      {
        rel: "stylesheet",
        href: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.10.0/css/all.min.css",
      },
      {
        rel: "stylesheet",
        href: "https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css",
      },
      // {
      //   rel: "stylesheet",
      //   href: "~/plugins/lib/owlcarousel/assets/owl.carousel.min.css",
      // },
      // {
      //   rel: "stylesheet",
      //   href: "~/plugins/lib/owlcarousel/assets/owl.theme.default.min.css",
      // },
    ],
    script: [
      {
        src: "https://code.jquery.com/jquery-3.4.1.min.js",
        type: "text/javascript",
      },
      {
        src: "https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js",
        type: "text/javascript",
      },
      {
        src: "https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js",
        type: "text/javascript",
      },
      // {
      //   src:"~/plugins/lib/owlcarousel/owl.carousel.min.js",
      //   type: "text/javascript",
      // },
      // {
      //   src: "~/plugins/lib/easing/easing.min.js",
      //   type: "text/javascript",
      // },
      // {
      //   src: "~/plugins/js/main.js",
      //   type: "text/javascript",
      // },
    ],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '~/assets/css/style.css'
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    // { src: "~/plugins/js/main.js", ssr: false },
    // { src: "~/plugins/lib/owlcarousel/owl.carousel.min.js", ssr: false },
    // { src: "~/plugins/lib/easing/easing.min.js", ssr: false },
    {src:'~/plugins/notification.js', ssr:false},

  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth-next'
  ],
  axios: {
    baseURL: 'http://localhost:8081', // Used as fallback if no runtime config is provided
  },
  auth: {
    strategies: {
      local: {
        user: {
          property: false,
          autoFetch: true
        },
        token: {
          property: 'token',
          required: true,
          type: 'Bearer'
        },
        endpoints: {
          login: { url: "/signin", method: "post" },
          logout: { url: "/logout", method: "delete" },
          user: { url: "/admin/user/me", method: "get"},
        },
        // redirect: {
        //   login: '/login',
        //   logout: '/login',
        //   callback: '/login',
        //   home: '/'
        // }
      },
    },
    redirect: {
      login: "/login",
      logout: "/login",
      callback: "/login",
      home: "/"
    }
  },
  router: {
    middleware: ['auth']
  },
  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    plugins: [
      new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "window.jQuery": "jquery",
      }),
    ],
  }
}

