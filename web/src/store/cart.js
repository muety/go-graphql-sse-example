import Vue from 'vue'

const state = () => ({
    products: []
})

const getters = {
    sum: (state) => {
        return state.products.reduce((acc, val) => acc += val.price, 0)
    },
    productIds: (state) => {
        return state.products.map(p => p.id)
    }
}

const mutations = {
    addProduct(state, product) {
        Vue.set(state, 'products', [...state.products, product])
    },
    clear(state) {
        Vue.set(state, 'products', [])
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
