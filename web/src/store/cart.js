import Vue from 'vue'
import {Cart} from '../model/cart'

const state = () => ({
    cart: new Cart({})
})

const getters = {
}

const mutations = {
    addProduct(state, product) {
        Vue.set(state.cart, 'products', [...state.cart.products, product])
    },
    clear(state) {
        Vue.set(state.cart, 'products', [])
    }
}

const actions = {}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
