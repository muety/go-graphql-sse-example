<template>
    <div class="container flex-col justify-center order-item">
        <div class="container justify-between items-center bg-brand padding-default">
            <h3 class="margin-r-default text-white">Order #{{ order.queueId }}</h3>
            <div class="badge critical margin-l-default">{{ displayTime }}</div>
        </div>
        <div class="padding-default">
            <div class="order-field text-sm container flex-col justify-center items-start margin-y-default">
                <div v-for="(p, i) in order.products" :key="i">
                    {{i+1}}. {{ p.name }}
                </div>
            </div>
            <div class="order-field text-sm container justify-center margin-y-default">
                {{ price(order.totalSum) }} €
            </div>
            <div class="order-field container justify-center margin-t-default" v-show="buttonText">
                <button @click="$emit('select', order.id)" v-html="buttonText"></button>
            </div>
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
            },
            buttonText() {
                switch (this.order.status) {
                    case 'pending':
                        return '&#10003; Done'
                    case 'delivering':
                        return '&#187; Delivered'
                    default:
                        return null
                }
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
        margin-right: 10px;
        margin-left: 10px;
        display: flex;
        align-items: center;
    }
</style>