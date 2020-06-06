# Database Initialization

* Create an empty MongoDB database (e.g. `test_db`)
* Enter the `mongo` shell and connect to your database
* Add a new user for the database:

```javascript
use test_db;

db.createUser({
    user: 'test_user',
    pwd: passwordPrompt(),
    roles: [
        { role: "readWrite", db: "test_db" }
    ]
});
```

* Create two required collections:
```javascript
use test_db;

db.createCollection('orders');
db.createCollection('products');
```

* Insert a few demo products:
```javascript
use test_db;

db.products.insert({
    name: "Schnitzel with fries",
    description: "Just a good, old, delicious Schnitzel with fries. Nothing more, nothing less.",
    price: 11.90
});

db.products.insert({
    name: "Sausages with potato salad",
    description: "In case you are a little less hungry.",
    price: 6.90
});
```