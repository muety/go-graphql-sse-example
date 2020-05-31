package service

import (
	eventhub "github.com/leandro-lugaresi/hub"
	"github.com/muety/go-graphql-sse-example/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface{}

type ServiceConfig struct {
	Hub *eventhub.Hub
	cfg *config.Config
	db  *mongo.Database
}
