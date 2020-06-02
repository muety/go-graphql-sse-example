<template>
    <div class="container justify-between padding-default margin-y-default order-item">
        <div class="order-field badge text-sm">
            # {{ order.queueId }}
        </div>
        <div class="order-field critical">
            {{ displayTime }} ago
        </div>
        <div class="order-field text-sm container flex-col justify-center items-start">
            <div v-for="(p, i) in order.products" :key="i">
                {{ p.name }}
            </div>
        </div>
        <div class="order-field text-sm">
            {{ price(order.totalSum) }} €
        </div>
        <div class="order-field">
            <button @click="$emit('select', order.id)">Done!️</button>
        </div>
    </div>
</template>

<script>
    import {Order} from '../model/order'
    import {addMinutes} from '../util/date'

    export default {
        name: 'OrderItem',
        props: {
            order: {
                type: Order,
                required: true
            }
        },
        data() {
            return {
                nowTime: new Date()
            }
        },
        computed: {
            displayTime() {
                if (addMinutes(this.order.createdAt, 59) > this.nowTime) return this.timeDiff(this.order.createdAt, true, true)
                return '∞'
            }
        },
        created() {
            setInterval(() => {
                this.nowTime = new Date()
            }, 5000)
        }
    }
</script>

<style scoped>
    .order-item {
        border: 1px solid #ddd;
    }

    .order-field {
        margin: 0 10px;
        display: flex;
        align-items: center;
    }

    .critical {
        color: #e49918;
        font-weight: bold;
    }
</style>