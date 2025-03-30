package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}, &Tag{})
	productRepository := NewProductRepository(db)
	// categoryRepository := NewCategoryRepository(db)

	// tags := []*Tag{
	// 	{Name: "PROMOTION"},
	// 	{Name: "NERD"},
	// 	{Name: "GEEK"},
	// 	{Name: "BESTSELLER"},
	// 	{Name: "NEW"},
	// }	
	// for _, t := range tags {
	// 	db.Create(t)
	// }

	// eletronics := &Category{Name: "Electronics"}
	// books := &Category{Name: "Books"}
	// categories := []*Category{
	// 	eletronics,
	// 	books,
	// }
	// for _, c := range categories {
	// 	categoryRepository.create(c)
	// }

	// products := []Product{
	// 	{Name: "Notebbok", Price: 1000, CategoryID: eletronics.ID, Tags: []Tag{*tags[1], *tags[2],  *tags[4]}},
	// 	{Name: "Mouse", Price: 30, CategoryID: eletronics.ID, Tags: []Tag{*tags[0], *tags[1], *tags[2]}},
	// 	{Name: "Game of Thrones", Price: 50, CategoryID: books.ID, Tags: []Tag{*tags[3], *tags[1], *tags[2]}},
	// }
	// for _, p := range products {
	// 	productRepository.create(&p)
	// }

	products2 := productRepository.findAllWithRelations()
	for _, p := range *products2 {
		fmt.Println("Product: ", p.Name)
		fmt.Println("Category: ", p.Category.Name)
		fmt.Println("Serial Number: ", p.SerialNumber.Number)
		fmt.Println("Tags: ")
		for _, t := range p.Tags {
			fmt.Println(" - ", t.Name)
		}
		fmt.Println()
	}

	var tags2 []Tag
	db.Preload("Products").Find(&tags2)
	for _, t := range tags2 {
		fmt.Printf("\nTag %s, Products:\n", t.Name)
		for _, p := range t.Products {
			fmt.Println("-", p.Name)
		}
	}
}