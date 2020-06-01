const imageUrlTpl = '/static/images/products/{id}.png'

class Product {
    constructor({id, name, description, price}) {
        this.id = id
        this.name = name
        this.description = description
        this.price = price
    }

    get imageUrl() {
        return imageUrlTpl.replace('{id}', this.id)
    }

    static new(data) {
        return new Product(data)
    }
}

export {
    Product,
}
