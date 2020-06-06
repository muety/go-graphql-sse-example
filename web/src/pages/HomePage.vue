<template>
    <div>
        <div class="container justify-center flex-col items-center padding-x-default">
            <div class="container flex-col justify-start items-start w-full constrained-container">
                <h2>Products</h2>
            </div>
            <div class="constrained-container w-full">
                <div v-for="p in products" :key="p.id">
                    <product-item :product="p" @select="onProductSelected"/>
                </div>
            </div>
        </div>
        <div class="cart container justify-center">
            <div class="container constrained-container padding-default justify-around w-full">
                <div class="container">
                    <h3 class="margin-r-default">🛒 {{ cart.products.length }} items</h3>
                    <span>({{ price(cart.sum) }} €)</span>
                </div>
                <div v-show="cart.products.length">
                    <button style="margin-right: 2px;" @click="clear">&#10226; Reset</button>
                    <button style="margin-left: 2px;" @click="checkout">&#10003; Checkout</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import ProductItem from '../components/product-item'
    import {mapActions, mapGetters, mapMutations, mapState} from 'vuex'

    export default {
        name: 'HomePage',
        components: {ProductItem},
        computed: {
            ...mapState('products', ['products']),
            ...mapState('cart', ['cart']),
            ...mapState('orders', ['myOrder']),
            ...mapGetters('products', ['product']),
        },
        methods: {
            ...mapActions('products', ['fetchProducts']),
            ...mapActions('orders', ['placeOrder']),
            ...mapMutations('cart', ['addProduct', 'clear']),
            onProductSelected(id) {
                this.addProduct(this.product(id))
            },
            checkout() {
                this.placeOrder({items: this.cart.productIds})
                    .then(() => this.$router.replace({name: 'order', params: {orderId: this.myOrder.id}}))
                    .catch(() => alert('Failed to place order'))
            }
        },
        created() {
            this.fetchProducts()
                .catch(() => alert('Failed to fetch products'))
        }
    }
</script>

<style scoped>
    .constrained-container {
        max-width: 600px;
    }

    .cart {
        position: absolute;
        bottom: 0;
        border-top: 2px solid #42b983;
        width: 100%;
        margin: 0;
    }
</style>