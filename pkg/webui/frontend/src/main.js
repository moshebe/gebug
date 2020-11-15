import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import Vuelidate from 'vuelidate'
import App from './App.vue'
import VueMaterial from 'vue-material'
// import VueFormulate from '@braid/vue-formulate'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

Vue.config.productionTip = false
Vue.use(Vuelidate)
Vue.use(VueMaterial)

// Vue.use(VueFormulate.default)

new Vue({
  render: h => h(App),
}).$mount('#app')
