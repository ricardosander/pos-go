package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	product := NewProduct("Mouse", 100.00)
	fmt.Println(product)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Name = "Mouse sem fio"
	product.Price = 150.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Produto:")
	fmt.Printf("Produc %s custa %.2f\n", p.Name, p.Price)

	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Products:")
	for _, p := range products {
		fmt.Printf("Product %s - %s costs %.2f\n", p.ID, p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}
