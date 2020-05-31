package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/muety/go-graphql-sse-example/model"
)

type productResolver struct {
	p *model.Product
}

func (r *productResolver) Id() graphql.ID {
	return graphql.ID(r.p.Id)
}

func (r *productResolver) Name() *string {
	return &r.p.Name
}

func (r *productResolver) Description() *string {
	return &r.p.Description
}

func (r *productResolver) Price() *float64 {
	return &r.p.Price
}
