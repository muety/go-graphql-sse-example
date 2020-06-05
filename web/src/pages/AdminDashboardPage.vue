<template>
    <div class="container justify-center flex-col items-center">
        <!-- Pending -->
        <div class="container flex-col justify-start items-start w-full constrained-container">
            <h2 class="margin-x-default">Pending Orders</h2>
        </div>
        <div class="container constrained-container w-full">
            <div v-if="!pendingOrders.length" class="margin-x-default text-left">No pending orders ...</div>
            <div v-for="o in pendingOrders" :key="o.id">
                <order-item :order="o" class="margin-default" @select="onOrderSelected($event, 'delivering')"/>
            </div>
        </div>

        <!-- Delivering -->
        <div v-if="deliveringOrders.length"
             class="container flex-col justify-start items-start w-full constrained-container margin-t-triple">
            <h2 class="margin-x-default">Delivering Orders</h2>
        </div>
        <div v-if="deliveringOrders.length" class="container constrained-container w-full">
            <div v-for="o in deliveringOrders" :key="o.id">
                <order-item :order="o" class="margin-default" @select="onOrderSelected($event, 'fulfilled')"/>
            </div>
        </div>

        <!-- Fulfilled -->
        <div v-if="fulfilledOrders.length"
             class="container flex-col justify-start items-start w-full constrained-container margin-t-triple">
            <h2 class="margin-x-default">Fulfilled Orders</h2>
        </div>
        <div v-if="fulfilledOrders.length" class="container constrained-container w-full">
            <div v-for="o in deliveringOrders" :key="o.id">
                <order-item :order="o" class="margin-default"/>
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
            ...mapActions('orders', ['fetchOrders', 'modifyOrder', 'subscribeOrderCreated']),
            onOrderSelected(id, nextState) {
                this.modifyOrder({id, status: nextState})
                    .catch(() => alert('Failed to update order'))
            }
        },
        created() {
            this.fetchOrders({status: 'pending', full: true, overwrite: true})
                .then(() => this.fetchOrders({status: 'delivering', full: true}))
                .then(() => this.fetchOrders({status: 'fulfilled', full: true}))

            this.subscribeOrderCreated()
                .catch(() => alert('Failed to subscribe to live updates'))
        }
    }
</script>

<style scoped>
    .constrained-container {
        max-width: 600px;
    }
</style>