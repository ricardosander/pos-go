package main

import "gorm.io/gorm"

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

type categoryRepository struct {
	db *gorm.DB
}

func (cr *categoryRepository) create(c *Category) {
	cr.db.Create(c)
}

func (cr *categoryRepository) findAll() []Category {
	var categories []Category
	err := cr.db.Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	return categories
}

func (cr *categoryRepository) findFirstByQuery(query string, value interface{}) Category {
	var category Category
	cr.db.Find(query, value).First(&category)
	return category
}