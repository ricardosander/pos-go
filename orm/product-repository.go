package main

import (
	"gorm.io/gorm"
)

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
		serialNumberRepository: NewSerialNumberRepository(db),
	}
}

type productRepository struct {
	db *gorm.DB
	serialNumberRepository *serialNumberRepository
}

func (p productRepository) create(product *Product) {
	p.db.Create(product)
	p.serialNumberRepository.create(&SerialNumber{ProductID: product.ID, Number: "123455"})
}

func (p productRepository) update(product *Product) {
	p.db.Save(product)
}

func (p productRepository) delete(product *Product) {
	p.db.Delete(product)
}

func (p productRepository) createBatch(products *[]Product) {
	p.db.Create(products)
}

func (p productRepository) firstById(id int) *Product {
	var product Product
	p.db.First(&product, id)
	return &product
}

func (p productRepository) findFirstByQuery(query string, value interface{}) *Product {
	var product Product
	p.db.First(&product, query, value)
	return &product
}

func (p productRepository) findAll() *[]Product {
	var products []Product
	p.db.Find(&products)
	return &products
}

func (p productRepository) findAllWithRelations() *[]Product {
	var products []Product
	p.db.Preload("Tags").Preload("Category").Preload("SerialNumber").Find(&products)
	return &products
}

func (p productRepository) limitQuery(limit int) *[]Product {
	var products []Product
	p.db.Limit(limit).Find(&products)
	return &products
}

func (p productRepository) limitOffetQuery(limit, offset int) *[]Product {
	var products []Product
	p.db.Limit(limit).Offset(offset).Find(&products)
	return &products
}

func (p productRepository) whereQuery(query string, value interface{}) *[]Product {
	var products []Product
	p.db.Where(query, value).Find(&products)
	return &products
}

func (p productRepository) likeQuery(likeQuery string) *[]Product {
	var products []Product
	p.db.Where("name LIKE ?", likeQuery).Find(&products)
	return &products
}
