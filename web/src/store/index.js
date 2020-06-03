import Vue from 'vue'
import Vuex from 'vuex'

import products from './products'
import orders from './orders'
import cart from './cart'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        products,
        orders,
        cart,
    },
    strict: debug
})
