import Vue from 'vue'
import VueRouter from 'vue-router'
import HomePage from './pages/HomePage'
import OrderStatusPage from './pages/OrderStatusPage'
import AdminDashboardPage from './pages/AdminDashboardPage'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomePage
    },
    {
        path: '/order/:orderId',
        name: 'order',
        component: OrderStatusPage
    },
    {
        path: '/admin',
        name: 'admin',
        component: AdminDashboardPage,
        beforeEnter: (to, from, next) => {
            // TODO: perform admin authorization here
            next()
        }
    }
]

const router = new VueRouter({
    routes
})

export default router
