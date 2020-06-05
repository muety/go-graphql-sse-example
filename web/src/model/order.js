import {addMinutes} from '../util/date'
import {Product} from './product'

const orderTimeoutMin = 10

class Order {
  constructor({id, queueId, createdAt, updatedAt, status, eta, totalSum, products}) {
    this.id = id
    this.queueId = queueId
    this.createdAt = new Date(createdAt)
    this.updatedAt = new Date(updatedAt)
    this.status = status
    this.eta = new Date(eta)
    this.totalSum = totalSum

    if (products) this.products = products.map(Product.new)
    else this.products = []
  }

  get isPending() {
    return this.status === 'pending'
  }

  get isDelivering() {
    return this.status === 'delivering'
  }

  get isFulfilled() {
    return this.status === 'fulfilled'
  }

  get isRejected() {
    return this.status === 'rejected'
  }

  get isTimedOut() {
    return addMinutes(new Date(), -orderTimeoutMin) >= this.createdAt
  }

  static new(data) {
    return new Order(data)
  }
}

export {
  Order
}
