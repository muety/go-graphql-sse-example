# go-graphql-sse-example

Basic example application to demonstrate an alternative way of building web APIs compared to REST. It combines these technologies:

* [Go](https://golang.org) as the primary backend-side language 
* [GraphQL](https://graphql.org/) as a "protocol" for defining web interfaces
* [Server-Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events) as a simple protocol for live updates, used as an implementation of [GraphQL Subscriptions](https://graphql.org/blog/subscriptions-in-graphql-and-relay/) here
* [MongoDB](https://www.mongodb.com/) as a flexible document database for storage
* [VueJS](https://vuejs.org/) as a frontend framework to build single-page web applications

This project is intended to serve as a starting point to build modern, GraphQL-based single-page web application with a clean, opinionated code structure in backend and frontend.

**[Original blog post](https://muetsch.io/modern-reactive-web-apis-with-graphql-go-and-server-sent-events-part-1.html)**

Inspired by: 
* [OscarYuen/go-graphql-starter](https://github.com/OscarYuen/go-graphql-starter)
* [andoshin11/clean-architecture-example-vue](https://github.com/andoshin11/clean-architecture-example-vue)

## Requirements
* Go >= 1.13
* NodeJS >= 12.8.x
* A MongoDB database
    * You can get a free one at [mlab.com](https://mlab.com/)

## Usage
### Backend (`/`)
* Edit `config.yml` to configure your database
* [Set up](docs/database.md) the database (i.e. add users, collections and demo data)
* Compile the GraphQL schema (located in [`schema`](schema)) to Go code: `./generate_schema.sh`
* Compile the server code: `GO111MODULE=on go build`
* Run the server: `./go-sse-graphql-example`

### Frontend (`/web`)
* Go to [`web`](web)
* Install dependencies: `yarn install`
* Run the frontend for development: `yarn serve`
* Go to http://localhost:4000 

The interactive GraphQL explorer can be accessed at http://localhost:8080/graphiql. You can find [example queries here](docs/queries.md).

## Limitations
* This example focuses primarily on the backend and the data flow between front- and backend. Accordingly, the UI is not very beautiful.
* Only very basic business logic is implemented, just enough to demonstrate the above concepts on a real-world example.

## Still to do
The current state of this project is a very basic working example. However, in order to use this as a base to build a fully-fledged application, there are still a few things missing, including the following.

* Authentication and authorization for GraphQL methods
* Caching and / or implementation of a [Dataloader](https://github.com/graphql/dataloader) for efficiency

Hopefully, I will have enough time to supplement these features in the future.

## License
MIT 