package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/muety/go-graphql-sse-example/model"
	"github.com/muety/go-graphql-sse-example/service"
	"time"
)

type orderResolver struct {
	o *model.Order
}

func (r *orderResolver) Id() graphql.ID {
	return graphql.ID(r.o.Id)
}

func (r *orderResolver) QueueId() *string {
	return &r.o.QueueId
}

func (r *orderResolver) CreatedAt() *graphql.Time {
	if r.o.CreatedAt != nil {
		return &graphql.Time{Time: *r.o.CreatedAt}
	}
	return nil
}

func (r *orderResolver) UpdatedAt() *graphql.Time {
	if r.o.CreatedAt != nil {
		return &graphql.Time{Time: *r.o.CreatedAt}
	}
	return nil
}

func (r *orderResolver) Eta() *graphql.Time {
	if r.o.Eta != nil {
		return &graphql.Time{Time: *r.o.Eta}
	}
	return nil
}

func (r *orderResolver) Status() *string {
	s := string(r.o.Status)
	return &s
}

func (r *orderResolver) TotalSum() *float64 {
	return &r.o.TotalSum
}

func (r *orderResolver) Products(ctx context.Context) (*[]*productResolver, error) {
	l := make([]*productResolver, len(r.o.Items))
	productIds := make([]string, len(r.o.Items))

	for i, item := range r.o.Items {
		productIds[i] = item
	}

	products, err := ctx.Value(service.KeyProductService).(*service.ProductService).GetBatch(productIds)
	if err != nil {
		return nil, err
	}

	for i, p := range products {
		l[i] = &productResolver{p: p}
	}

	return &l, nil
}

// additional types

type OrderInput struct {
	Items []*string
}

func (i *OrderInput) ToEntity() *model.Order {
	items := make([]string, len(i.Items))

	for i, it := range i.Items {
		items[i] = *it
	}

	return &model.Order{
		Items: items,
	}
}

type OrderUpdateInput struct {
	Id     string
	Status *model.OrderStatus
	Eta    *time.Time
}

func (i *OrderUpdateInput) ToEntity() *model.Order {
	o := &model.Order{
		Id:  i.Id,
		Eta: i.Eta,
	}
	if i.Status != nil {
		o.Status = *i.Status
	}

	return o
}
