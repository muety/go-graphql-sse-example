<template>
    <div class="container justify-center flex-col items-center">
        <!-- Pending -->
        <div class="container flex-col justify-start items-start w-full constrained-container">
            <h2>Pending Orders</h2>
        </div>
        <div class="constrained-container w-full">
            <div v-if="!pendingOrders.length" class="text-left">No pending orders ...</div>
            <div v-for="o in pendingOrders" :key="o.id">
                <order-item :order="o"/>
            </div>
        </div>

        <!-- Delivering -->
        <div v-if="deliveringOrders.length"
             class="container flex-col justify-start items-start w-full constrained-container margin-t-triple">
            <h2>Delivering Orders</h2>
        </div>
        <div v-if="deliveringOrders.length" class="constrained-container w-full">
            <div v-for="o in deliveringOrders" :key="o.id">
                <order-item :order="o"/>
            </div>
        </div>

        <!-- Fulfilled -->
        <div v-if="fulfilledOrders.length"
             class="container flex-col justify-start items-start w-full constrained-container margin-t-triple">
            <h2>Fulfilled Orders</h2>
        </div>
        <div v-if="fulfilledOrders.length" class="constrained-container w-full">
            <div v-for="o in deliveringOrders" :key="o.id">
                <order-item :order="o"/>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapActions, mapGetters} from 'vuex'
    import OrderItem from '../components/order-item'

    export default {
        name: 'AdminDashboardPage',
        components: {OrderItem},
        computed: {
            ...mapGetters('orders', ['pendingOrders', 'deliveringOrders', 'fulfilledOrders'])
        },
        methods: {
            ...mapActions('orders', ['fetchOrders'])
        },
        created() {
            this.fetchOrders({status: 'pending', full: true, overwrite: true})
                .then(() => this.fetchOrders({status: 'delivering', full: true}))
                .then(() => this.fetchOrders({status: 'fulfilled', full: true}))
                .catch(() => alert('Failed to fetch orders'))
        }
    }
</script>

<style scoped>
    .constrained-container {
        max-width: 600px;
    }
</style>