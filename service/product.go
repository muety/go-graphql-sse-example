package service

import (
	"context"
	"github.com/muety/go-graphql-sse-example/config"
	"github.com/muety/go-graphql-sse-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	*ServiceConfig
}

const (
	KeyProductService = "productService"
)

func NewProductService() *ProductService {

	return &ProductService{&ServiceConfig{
		config.GetEventHub(),
		config.Get(),
		config.GetDb(),
	}}
}

func (s *ProductService) Get(id string) (product *model.Product, err error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	if err = config.GetDbCollection(config.CollectionProducts).FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&product); err != nil {
		return product, err
	}
	return product, err
}

func (s *ProductService) GetBatch(ids []string) (products []*model.Product, err error) {
	oids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		oid, _ := primitive.ObjectIDFromHex(id)
		oids[i] = oid
	}
	filter := bson.M{"_id": bson.M{"$in": oids}}

	cur, err := config.GetDbCollection(config.CollectionProducts).Find(context.TODO(), filter)
	if err != nil {
		return products, err
	}

	if err := cur.All(context.TODO(), &products); err != nil {
		return products, err
	}

	return products, err
}

func (s *ProductService) GetBatchMap(ids []string) (productMap map[string]*model.Product, err error) {
	productMap = make(map[string]*model.Product)
	products, err := s.GetBatch(ids)
	if err != nil {
		return productMap, err
	}

	for _, p := range products {
		productMap[p.Id] = p
	}

	return productMap, err
}

func (s *ProductService) GetAll() (products []*model.Product, err error) {
	cur, err := config.GetDbCollection(config.CollectionProducts).Find(context.TODO(), bson.D{})
	if err != nil {
		return products, err
	}

	if err := cur.All(context.TODO(), &products); err != nil {
		return products, err
	}

	return products, err
}
