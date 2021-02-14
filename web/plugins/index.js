import define from './define'
import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default (ctx, inject) => {
  inject('define', define)
  inject('isMd', () => ctx.$vuetify.breakpoint.mdAndUp)
}