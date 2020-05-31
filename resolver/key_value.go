package resolver

import (
	"fmt"
	"github.com/muety/go-graphql-sse-example/model"
)

type keyValueResolver struct {
	a *model.KeyValue
}

func (r *keyValueResolver) Key() *string {
	return &r.a.Key
}

func (r *keyValueResolver) Value() *string {
	s := fmt.Sprintf("%v", r.a.Value)
	return &s
}

// additional types

type KeyValueInput struct {
	Key   string
	Value string
}

func (i *KeyValueInput) ToEntity() *model.KeyValue {
	return &model.KeyValue{
		Key:   i.Key,
		Value: i.Value,
	}
}
