import Vue from 'vue'
import {Order} from '../model/order'

let activeEventSources = []

const orderSelection = `
  id
  queueId
  createdAt
  updatedAt
  status
  eta
  totalSum
`

const orderSelectionFull = `
  ${orderSelection}
  products {
    id
    name
  }
`

const state = () => ({
    orders: [],
    myOrder: null
})

const getters = {
    orderRefs: (state) => (id) => {
        return state.orders.filter(o => o.id === id)
            .concat(state.myOrder && state.myOrder.id === id ? [state.myOrder] : [])
    },
    pendingOrders: (state) => {
        return state.orders.filter(o => o.isPending)
    },
    deliveringOrders: (state) => {
        return state.orders.filter(o => o.isDelivering)
    },
    fulfilledOrders: (state) => {
        return state.orders.filter(o => o.isFulfilled)
    },
}

const mutations = {
    setMyOrder(state, order) {
        Vue.set(state, 'myOrder', order)
    },
    addOrder(state, order) {
        state.orders.push(order)
    },
    setOrders(state, orders) {
        Vue.set(state, 'orders', orders)
    },
    addOrders(state, orders) {
        Vue.set(state, 'orders', state.orders.concat(orders))
    },
    updateOrder(state, update) {
        // little hack to have this store handle orders for both admins and customers
        const orderRefs = state.orders.filter(o => o.id === update.id)
            .concat(state.myOrder && state.myOrder.id === update.id ? [state.myOrder] : [])

        orderRefs.forEach(ref => {
            Object.entries(update).forEach(e => {
                Vue.set(ref, e[0], e[1])
            })
        })
    },
    clearSubscriptions() {
        activeEventSources.forEach(s => s.disconnect())
        activeEventSources = []
    }
}

const actions = {
    async fetchOrder({commit}, {id, full, target}) {
        const q = `query {
          order(id: "${id}") {
            ${full ? orderSelectionFull : orderSelection}
          }
        }`

        const data = await Vue.$api.graphql.request(q, null)

        if (target === 'orders') {
            commit('addOrder', new Order(data.order))
        } else {
            commit('setMyOrder', new Order(data.order))
        }
    },

    async fetchOrders({commit}, {status, full, overwrite}) {
        const q = `query {
          orders(status: "${status}") {
            ${full ? orderSelectionFull : orderSelection}
          }
        }`

        const data = await Vue.$api.graphql.request(q, null)

        const commitFunc = overwrite ? 'setOrders' : 'addOrders'
        commit(commitFunc, data.orders.map(Order.new))
    },

    async placeOrder({commit}, cart) {
        const q = `mutation($order: OrderInput!) {
          createOrder(order: $order) {
            id
            createdAt
            updatedAt
            eta
            status
            totalSum
            queueId
          }
        }`

        // OrderInput ~ Cart
        const vars = {
            order: cart
        }

        const data = await Vue.$api.graphql.request(q, vars)
        commit('setMyOrder', new Order(data.createOrder))
    },

    async modifyOrder({commit}, {id, status, eta}) {
        const q = `mutation($order: OrderUpdateInput!) {
          updateOrder(order: $order) {
            id
            status
            eta
          }
        }`

        const order = {id}
        if (status) order.status = status
        if (eta) order.eta = eta

        const vars = {
            order
        }

        await Vue.$api.graphql.request(q, vars)
        commit('updateOrder', order)
    },

    async subscribeOrderChanged({commit, getters}, {id}) {
        const orderRefs = getters.orderRefs(id)
        if (!orderRefs.length) return

        const q = `subscription {
            orderChanged(id: "${id}") {
                ${orderSelection}
           }
        }`

        const source = Vue.$api.sse.request(q, null)
        source.addEventListener('message', e => {
            const payload = JSON.parse(e.data)
            if (!payload || !payload.data || !payload.data.orderChanged) {
                console.error('property "orderChanged" not present on update event')
                return
            }
            commit('updateOrder', new Order(payload.data.orderChanged))
        })

        activeEventSources.push(source)
        source.stream()
    },

    async subscribeOrderCreated({commit}) {
        const q = `subscription {
            orderCreated() {
                ${orderSelectionFull}
           }
        }`

        const source = Vue.$api.sse.request(q, null)
        source.addEventListener('message', e => {
            const payload = JSON.parse(e.data)
            if (!payload || !payload.data || !payload.data.orderCreated) {
                console.error('property "orderCreated" not present on update event')
                return
            }
            commit('addOrder', new Order(payload.data.orderCreated))
        })

        activeEventSources.push(source)
        source.stream()
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
