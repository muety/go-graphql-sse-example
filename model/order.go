package model

import (
	"time"
)

type OrderStatus string

const (
	KeyOrder              = "order"
	OrderStatusPending    = "pending"
	OrderStatusDelivering = "delivering"
	OrderStatusFulfilled  = "fulfilled"
	OrderStatusRejected   = "rejected"
)

type Order struct {
	Id        string     `bson:"_id,omitempty"`
	QueueId   string     `bson:"queueId"`
	CreatedAt *time.Time `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
	Status    OrderStatus
	Eta       *time.Time
	TotalSum  float64  `bson:"totalSum"`
	Items     []string // product ids
}

type OrderItem struct {
	Id         string `bson:"_id,omitempty"`
	Attributes []*KeyValue
}

func (o *Order) CalcTotalSum(resolve func(pid string) (*Product, error)) (float64, error) {
	var sum float64

	for _, it := range o.Items {
		p, err := resolve(it)
		if err != nil {
			return 0, err
		}

		sum += p.Price
	}

	return sum, nil
}
