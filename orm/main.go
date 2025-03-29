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
	dao := NewProductDTO(db)
	
	dao.create(&Product{
		Name:  "Product A",
		Price: 100.0,
	})

	products := &[]Product{
		{Name: "Product B", Price: 200.0},
		{Name: "Product C", Price: 300.0},
	}
	dao.createBatch(products)

	product := dao.firstById(1)
	fmt.Println("Product by id:")
	product.print()

	product = dao.findFirstByQuery("name = ?", "Product B")
	fmt.Println("Product by query:")
	product.print()

	products = dao.findAll()
	fmt.Println("All products:")
	print(*products)

	products = dao.limitQuery(2)
	fmt.Println("Limit query:")
	print(*products)

	products = dao.limitOffetQuery(2, 10)
	fmt.Println("Limit and offset query:")
	print(*products)

	products = dao.whereQuery("price > ?", 150.0)
	fmt.Println("Where query:")
	print(*products)

	products = dao.likeQuery("% C%")
	fmt.Println("Like query:")
	print(*products)
	for _, p := range *products {
		p.Name = "Mouse"
		dao.update(&p)
	}

	products = dao.whereQuery("name = ?", "Mouse")
	fmt.Println("Updated products:")	
	print(*products)

	products = dao.whereQuery("name = ?", "Product B")
	for _, p := range *products {
		dao.delete(&p)
	}

	products = dao.findAll()
	fmt.Println("All products")
	print(*products)
}

func print(ps []Product) {
	for _, p := range ps {
		p.print()
	}
}