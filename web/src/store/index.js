import Vue from 'vue'
import Vuex from 'vuex'

import products from './products'
import orders from './orders'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        products,
        orders,
    },
    strict: debug
})
