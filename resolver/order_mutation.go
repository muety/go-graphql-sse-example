package resolver

import (
	"github.com/muety/go-graphql-sse-example/service"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateOrder(ctx context.Context, args struct{ Order OrderInput }) (*orderResolver, error) {
	order, err := ctx.Value(service.KeyOrderService).(*service.OrderService).Create(args.Order.ToEntity())
	if err != nil {
		return nil, err
	}

	return &orderResolver{order}, nil
}

func (r *Resolver) UpdateOrder(ctx context.Context, args struct{ Order OrderUpdateInput }) (*orderResolver, error) {
	order, err := ctx.Value(service.KeyOrderService).(*service.OrderService).Get(args.Order.Id)
	if err != nil {
		return nil, err
	}

	order, err = ctx.Value(service.KeyOrderService).(*service.OrderService).Update(args.Order.ToEntity())
	if err != nil {
		return nil, err
	}

	return &orderResolver{order}, nil
}
