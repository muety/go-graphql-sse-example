package service

import (
	"context"
	"errors"
	"fmt"
	eventhub "github.com/leandro-lugaresi/hub"
	"github.com/muety/go-graphql-sse-example/config"
	"github.com/muety/go-graphql-sse-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type OrderService struct {
	*ServiceConfig
	*ProductService
}

var lastId int

const (
	KeyOrderService    = "orderService"
	KeyOrderCreated    = "order.create"
	KeyTplOrderChanged = "order.update.%s"
)

func NewOrderService() *OrderService {
	return &OrderService{
		&ServiceConfig{
			config.GetEventHub(),
			config.Get(),
			config.GetDb(),
		},
		NewProductService(),
	}
}

func (s *OrderService) Get(id string) (order *model.Order, err error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	if err = config.GetDbCollection(config.CollectionOrders).FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&order); err != nil {
		return order, err
	}
	return order, err
}

func (s *OrderService) FindByStatus(status string) (orders []*model.Order, err error) {
	findOpts := options.Find()
	findOpts.SetSort(bson.D{{"createdAt", 1}})
	filter := bson.D{{"status", status}}

	cur, err := config.GetDbCollection(config.CollectionOrders).Find(context.TODO(), filter, findOpts)
	if err != nil {
		return orders, err
	}

	if err := cur.All(context.TODO(), &orders); err != nil {
		return orders, err
	}

	return orders, err
}

func (s *OrderService) Create(order *model.Order) (*model.Order, error) {
	if err := s.preprocessNew(order); err != nil {
		return nil, err
	}

	res, err := config.GetDbCollection(config.CollectionOrders).InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}

	order.Id = res.InsertedID.(primitive.ObjectID).Hex()
	s.Hub.Publish(eventhub.Message{
		Name:   KeyOrderCreated,
		Fields: eventhub.Fields{"id": order.Id},
	})

	return order, err
}

func (s *OrderService) Update(order *model.Order) (*model.Order, error) {
	oid, _ := primitive.ObjectIDFromHex(order.Id)
	update, err := s.processUpdate(order)
	if err != nil {
		return nil, err
	}

	res, err := config.GetDbCollection(config.CollectionOrders).UpdateOne(context.TODO(), bson.D{{"_id", oid}}, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, errors.New(fmt.Sprintf("no %s document matched id %s\n", model.KeyOrder, order.Id))
	}

	s.Hub.Publish(eventhub.Message{
		Name:   fmt.Sprintf(KeyTplOrderChanged, order.Id),
		Fields: eventhub.Fields{"id": order.Id},
	})

	return order, err
}

// Inplace!
func (s *OrderService) preprocessNew(order *model.Order) error {
	now := time.Now()

	// TODO: generate from database, to maintain the sequence between server restarts
	qid := strconv.Itoa(lastId)
	lastId++

	// also validates attributes and options implicitly
	sum, err := order.CalcTotalSum(s.ProductService.Get)
	if err != nil {
		return err
	}

	// default eta for new orders
	// TODO: make configurable
	eta := now.Add(10 * time.Minute)

	order.QueueId = qid
	order.CreatedAt = &now
	order.UpdatedAt = &now
	order.Status = model.OrderStatusPending
	order.Eta = &eta
	order.TotalSum = sum

	return nil
}

// Inplace!
func (s *OrderService) processUpdate(order *model.Order) (*bson.D, error) {
	data := bson.M{}
	now := time.Now()

	order.UpdatedAt = &now

	data["updatedAt"] = order.UpdatedAt
	data["status"] = order.Status
	data["eta"] = order.Eta

	return &bson.D{{"$set", data}}, nil
}
