# Running GraphQL queries

Once the server is up and running, you can access the interactive GraphQL browser at http://localhost:8080/graphiql.

## Example queries

### List products

```graphql
query {
  product {
    id
    name
    price
  }
}
```

### List pending orders
```graphql
query {
  orders(status: "pending") {
    id
    queueId
    createdAt
    products {
        name
    }
  }
}
```

### Create new order
```graphql
mutation($order: OrderInput!) {
  createOrder(order: $order) {
    id
    queueId
  }
}
```

With variables:
```json
{
  "order": {
    "items": [ "5eaaa77fcbfe964301372966" ]
  }
}
```