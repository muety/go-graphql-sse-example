import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import apiPlugin from './plugins/api'
import priceFilter from './filters/price'
import timeDiffFilter from './filters/time_diff'

import './styles/index.css'

Vue.config.productionTip = false

Vue.use(apiPlugin)

Vue.prototype.price = priceFilter
Vue.prototype.timeDiff = timeDiffFilter

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
