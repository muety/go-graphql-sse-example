<template>
    <div class="container justify-between padding-default margin-y-default product-item">
        <div class="product-field font-semibold">
            {{ product.name }}
        </div>
        <div class="product-field text-sm overflow-hidden">
            <p class="text-truncate" :title="product.description">{{ product.description }}</p>
        </div>
        <div class="product-field">
            {{ price(product.price) }} €
        </div>
        <div class="product-field">
            <button @click="$emit('select', product.id)">Purchase</button>
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
        display: flex;
        align-items: center;
        white-space: nowrap;
    }
</style>