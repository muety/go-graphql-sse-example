<template>
    <div class="container justify-center flex-col items-center">
        <div class="container flex-col justify-start items-start w-full constrained-container">
            <h2>Products</h2>
        </div>
        <div class="constrained-container w-full">
            <div v-for="p in products" :key="p.id">
                <product-item :product="p"/>
            </div>
        </div>
    </div>
</template>

<script>
    import ProductItem from '../components/product-item'
    import {mapActions, mapState} from 'vuex'

    export default {
        name: 'HomePage',
        components: {ProductItem},
        computed: {
            ...mapState('products', ['products'])
        },
        methods: {
            ...mapActions('products', ['fetchProducts'])
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
</style>