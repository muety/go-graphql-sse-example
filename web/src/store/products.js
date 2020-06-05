import Vue from 'vue'
import {Product} from '../model/product'

const state = () => ({
    products: []
})

const getters = {
    product: (state) => (id) => {
        return state.products.find(p => p.id === id)
    }
}

const mutations = {
    addProducts(state, products) {
        Vue.set(state, 'products', products)
    },
    addProduct(state, product) {
        const idx = state.products.findIndex(p => p.id === product.id)

        if (idx >= 0) {
            Vue.set(state.products, idx, product)
        } else {
            Vue.set(state, 'products', [...state.products, product])
        }
    }
}

const actions = {
    async fetchProducts({commit}) {
        const q = `{
          products() {
            id
            name
            description
            price
          }
        }`

        const data = await Vue.$api.graphql.request(q)
        commit('addProducts', data.products.map(Product.new))
    },
    async fetchProduct({commit}, id) {
        const q = `{
          product(id: "${id}") {
            id
            name
            description
            price
            order
          }
        }`

        const data = await Vue.$api.graphql.request(q)
        commit('addProduct', new Product(data.product))
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
