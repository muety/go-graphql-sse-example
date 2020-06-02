<template>
    <div class="container justify-between padding-default margin-y-default product-item">
        <div class="product-field">
            <span class="font-semibold">{{ product.name }}</span>
        </div>
        <div class="product-field grow">
            <span class="text-sm">{{ product.description }}</span>
        </div>
        <div class="product-field">
            <span class="">{{ price(product.price) }} €</span>
        </div>
        <div class="product-field">
            <button>Buy</button>
        </div>
    </div>
</template>

<script>
    import {Product} from '../model/product'

    export default {
        name: 'ProductItem',
        props: {
            product: {
                type: Product,
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
    .product-item {
        border: 1px solid #ddd;
    }

    .product-field {
        margin: 0 10px;
    }
</style>