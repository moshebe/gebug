import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'

import App from './App.vue'
// TODO:
// @ts-ignore
import Vuelidate from 'vuelidate'
// @ts-ignore
import VueMaterial from 'vue-material'
// @ts-ignore
import VueFormulate from '@braid/vue-formulate'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'


Vue.config.productionTip = false
Vue.use(Vuelidate)
Vue.use(VueMaterial)
Vue.use(VueFormulate)


new Vue({
  render: h => h(App),
}).$mount('#app')
