<template>
    <div class="container justify-between padding-default margin-y-default order-item">
        <div class="order-field badge text-sm">
            {{ order.queueId }}
        </div>
        <div class="order-field text-sm grow container flex-col justify-center items-start">
            <div v-for="(p, i) in order.products" :key="i">
                {{ p.name }}
            </div>
        </div>
    </div>
</template>

<script>
    import {Order} from '../model/order'

    export default {
        name: 'OrderItem',
        props: {
            order: {
                type: Order,
                required: true
            }
        },
        data() {
            return {}
        },
        computed: {}
    }
</script>

<style scoped>
    .order-item {
        border: 1px solid #ddd;
    }

    .order-field {
        margin: 0 10px;
    }
</style>