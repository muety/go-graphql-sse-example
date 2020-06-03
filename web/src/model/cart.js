class Cart {
    constructor ({ items }) {
        this.items = items || []
    }

    static new (data) {
        return new Cart(data)
    }
}

export {
    Cart
}