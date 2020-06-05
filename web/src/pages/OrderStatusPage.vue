<template>
    <div v-if="myOrder" class="container flex-col padding-quadruple">
        <progress-ring
                :stroke="4"
                :progress="progress"
                color="#42b983">
            <div class="container flex-col padding-default text-sm label-container">
                <span v-if="myOrder.isPending" class="text-bold">Your order is being processed.</span>
                <span v-else-if="myOrder.isDelivering" class="text-bold">Your order is ready, come pick it up!</span>
                <span v-else-if="myOrder.isFulfilled" class="text-bold">Your order is done.</span>
                <div v-if="myOrder.isPending || myOrder.isDelivering || myOrder.isRejected" class="container flex-col padding-t-default">
                    <span>Your pickup code is:</span>
                    <span class="text-bold text-4xl padding-t-default">{{ myOrder.queueId }}</span>

                    <span v-if="myOrder.isPending" class="padding-t-default">Expected waiting time:</span>
                    <span v-if="myOrder.isPending" class="text-bold text-lg">{{ timeDiff(myOrder.eta) }}</span>
                </div>
            </div>
        </progress-ring>
    </div>
</template>

<script>
    import {mapActions, mapMutations, mapState} from 'vuex'
    import ProgressRing from '../components/progress-ring'

    export default {
        name: 'OrderStatusPage',
        components: {
            ProgressRing
        },
        data() {
            return {
                orderId: '',
                nowTime: new Date()
            }
        },
        computed: {
            ...mapState('orders', ['myOrder']),
            progress() {
                const totalDiff = Math.max(this.myOrder.eta - this.myOrder.createdAt, 1)
                const currentDiff = Math.max(this.myOrder.eta - this.nowTime, 1)
                return (totalDiff - currentDiff) / totalDiff * 100
            }
        },
        created() {
            this.orderId = this.$route.params.orderId

            this.fetchOrder({id: this.orderId, full: false})
                .then(() => this.subscribeOrderChanged({id: this.orderId}))
                .catch(() => alert('Failed to fetch order'))

            setInterval(() => {
                this.nowTime = new Date()
            }, 3000)
        },
        destroyed() {
            this.clearSubscriptions()
        },
        methods: {
            ...mapActions('orders', ['fetchOrder', 'subscribeOrderChanged']),
            ...mapMutations('orders', ['clearSubscriptions'])
        }
    }
</script>

<style scoped>
    .label-container {
        background-color: #42b983;
        color: #ffffff;
        display: inline-block;
        box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06)
    }
</style>