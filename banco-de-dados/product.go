package main

import "github.com/google/uuid"

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

type Product struct {
	ID    string
	Name  string
	Price float64
}
