package main

import "fmt"

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	Products []Product
}

func (c Category) Print() {
	fmt.Printf("Category: %s\n", c.Name)
	for _, p := range c.Products {
		p.print()
	}
	fmt.Print("------\n")
}
