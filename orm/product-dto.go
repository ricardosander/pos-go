package main

import (
	"gorm.io/gorm"
)

func NewProductDTO(db *gorm.DB) *productDTO {
	newDB := productDTO{db: db}
	newDB.db.AutoMigrate(&Product{})
	return & newDB
}

type productDTO struct {
	db *gorm.DB
}

func (pd productDTO) create(product *Product) {
	pd.db.Create(&product)
}

func (pd productDTO) update(product *Product) {
	pd.db.Save(&product)
}

func (pd productDTO) delete(product *Product) {
	pd.db.Delete(product)
}

func (pd productDTO) createBatch(products *[]Product) {
	pd.db.Create(&products)
}

func (pd productDTO) firstById(id int) *Product {
	var p Product
	pd.db.First(&p, id)
	return &p
}

func (pd productDTO) findFirstByQuery(query string, value interface{}) *Product {
	var p Product
	pd.db.First(&p, query, value)
	return &p
}

func (pd productDTO) findAll() *[]Product {
	var ps []Product
	pd.db.Find(&ps)
	return &ps
}

func (pd productDTO) limitQuery(limit int) *[]Product {
	var ps []Product
	pd.db.Limit(limit).Find(&ps)
	return &ps
}

func (pd productDTO) limitOffetQuery(limit, offset int) *[]Product {
	var ps []Product
	pd.db.Limit(limit).Offset(offset).Find(&ps)
	return &ps
}

func (pd productDTO) whereQuery(query string, value interface{}) *[]Product {
	var ps []Product
	pd.db.Where(query, value).Find(&ps)
	return &ps
}

func (pd productDTO) likeQuery(likeQuery string) *[]Product {
	var ps []Product
	pd.db.Where("name LIKE ?", likeQuery).Find(&ps)
	return &ps
}
