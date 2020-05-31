package main

import (
	"github.com/graph-gophers/graphql-go"
	eventhub "github.com/leandro-lugaresi/hub"
	"github.com/muety/go-graphql-sse-example/config"
	"github.com/muety/go-graphql-sse-example/middleware"
	"github.com/muety/go-graphql-sse-example/resolver"
	"github.com/muety/go-graphql-sse-example/schema"
	"github.com/muety/go-graphql-sse-example/service"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

const KeyAppInit = "application.initialized"

func main() {
	cfg := config.Init("config.yml")
	ctx := context.Background()
	hub := config.InitEventHub()

	_, err := config.InitDbSync()
	if err != nil {
		log.Printf("error: unable to connect to db: %s \n", err)
	}

	productService := service.NewProductService()
	orderService := service.NewOrderService()

	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, service.KeyProductService, productService)
	ctx = context.WithValue(ctx, service.KeyOrderService, orderService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	http.Handle("/graphiql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/graphiql.html")
	}))

	// API routes
	http.Handle("/api/query", middleware.AddContext(ctx, &middleware.GraphQL{Schema: graphqlSchema}))

	hub.Publish(eventhub.Message{Name: KeyAppInit})

	log.Printf("listening on %s\n", cfg.ListenAddr())
	log.Println(http.ListenAndServe(cfg.ListenAddr(), nil))
}
