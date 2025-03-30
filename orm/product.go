package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Tag struct {
	ID  int `gorm:"primaryKey"`
	Name string
	Products []Product `gorm:"many2many:product_tags;"`
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string 
	Price float64
	SerialNumber SerialNumber
	CategoryID int
	Category Category
	Tags []Tag `gorm:"many2many:product_tags;"`
	gorm.Model
}

func (p Product) print() {
	fmt.Printf("Product: %s - SerialNumber %s - Category %s\n", p.Name, p.SerialNumber.Number, p.Category.Name)
}