import define from './define'
// import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import Vue from 'vue'
import Vuetify from 'vuetify/lib'


Vue.use(Vuetify)

export default (ctx, inject) => {
  inject('define', define)
  inject('isMd', () => ctx.$vuetify.breakpoint.mdAndUp)
}