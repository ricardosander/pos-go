package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string 
	Price float64
	gorm.Model
}

func (p Product) print() {
	fmt.Printf("Product: %v\n", p)
}