<template>
    <div>
        <div class="container justify-center flex-col items-center">
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
                    <h3 class="margin-r-default">🛒 {{ cart.length }} items</h3>
                    <span>({{ price(sum) }} €)</span>
                </div>
                <div>
                    <button style="margin-right: 2px;" @click="clear">&#10226; Reset</button>
                    <button style="margin-left: 2px;">&#10003; Checkout</button>
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
            ...mapState('cart', {cart: 'products'}),
            ...mapGetters('products', ['product']),
            ...mapGetters('cart', ['sum'])
        },
        methods: {
            ...mapActions('products', ['fetchProducts']),
            ...mapMutations('cart', ['addProduct', 'clear']),
            onProductSelected(id) {
                this.addProduct(this.product(id))
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