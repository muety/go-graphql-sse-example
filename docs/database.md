# Database Initialization

## Create a new database (e.g. `test_db`)
### Option 1: Hosted
Get a hosted database at [mlab](https://mlab.com)

### Option 2: Docker
```
docker run -d \
    --name mongo \
    --net mongo \
    -p 27017:27017 \
    -e "MONGO_INITDB_ROOT_USERNAME=admin" \
    -e "MONGO_INITDB_ROOT_PASSWORD=admin" \
    mongo
```

```
docker run -d \
    --name mongo-express \
    --net mongo \
    -p 8081:8081 \
    -e ME_CONFIG_MONGODB_SERVER=mongo \
    -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin \ 
    -e ME_CONFIG_MONGODB_ADMINPASSWORD=admin \ 
    mongo-express
```

You can now access a nice Mongo web UI at http://localhost:8081.

* Enter the `mongo` shell and connect to your database
```bash
mongo -u admin --authenticationDatabase admin
```

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

## Initialize collections

```javascript
use test_db;

db.createCollection('orders');
db.createCollection('products');
```

### Initialize demo data

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