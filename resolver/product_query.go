package resolver

import (
	"github.com/muety/go-graphql-sse-example/service"
	"golang.org/x/net/context"
)

func (r *Resolver) Product(ctx context.Context, args struct {
	Id string
}) (*productResolver, error) {
	product, err := ctx.Value(service.KeyProductService).(*service.ProductService).Get(args.Id)
	if err != nil {
		return nil, err
	}
	return &productResolver{product}, nil
}

func (r *Resolver) Products(ctx context.Context) (*[]*productResolver, error) {
	products, err := ctx.Value(service.KeyProductService).(*service.ProductService).GetAll()
	if err != nil {
		return nil, err
	}

	l := make([]*productResolver, len(products))
	for i := range l {
		l[i] = &productResolver{p: products[i]}
	}

	return &l, nil
}
