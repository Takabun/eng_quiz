// import colors from "vuetify/es5/util/colors";
// import { NuxtConfiguration }  from '@nuxt/config'
// import { NuxtConfig } from '@nuxt/types'

const nuxtConfig = {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: "universal",
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  env: {
    // todoApiUrl: process.env.TODO_API_URL || 'http://localhost:1323'
    todoApiUrl: process.env.TODO_API_URL ? `http://${process.env.TODO_API_URL}` : 'http://localhost:1323'
  },

  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: "%s - " + process.env.npm_package_name,
    title: process.env.npm_package_name || "",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: process.env.npm_package_description || ""
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [
    { src: 'plugins/index', ssr: false },
    { src: 'plugins/APICall', ssr: false }
  ],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    "@nuxtjs/vuetify",
    "@nuxt/typescript-build",
    "nuxt-typed-vuex"
  ],
  /*
   ** Nuxt.js modules
   */
  modules: ['@nuxtjs/proxy'],
  proxy: {
    '/awss3': {
      target: 'https://image.s3.amazonaws.com/',
      // pathRewrite: {
      //   '^/awss3' : '/'
      //   }
      }
  },
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ["~/assets/variables.scss"],
    theme: {
      dark: false,
      themes: {
        dark: {
          primary: 'red',
          accent: 'red',
          secondary: 'red',
          info: 'red',
          warning: 'red',
          error: 'red',
          success: 'red'
        }
      }
    }
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {
    transpile: [/typed-vuex/]
  },
  typescript: {
    typeCheck: true,
    ignoreNotFoundWarnings: true
  },
};
module.exports = nuxtConfig