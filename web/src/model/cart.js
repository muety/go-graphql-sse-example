import {Product} from './product'

class Cart {
    constructor({items: products}) {
        if (products) this.products = products.map(Product.new)
        else this.products = []
    }

    get sum() {
        return this.products.reduce((acc, val) => acc + val.price, 0)
    }

    get productIds() {
        return this.products.map(p => p.id)
    }

    static new(data) {
        return new Cart(data)
    }
}

export {
    Cart
}