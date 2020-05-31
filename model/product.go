package model

type Product struct {
	Id          string `bson:"_id,omitempty"`
	Name        string
	Description string
	Price       float64
}
