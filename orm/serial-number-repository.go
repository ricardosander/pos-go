package main

import "gorm.io/gorm"

func NewSerialNumberRepository(db *gorm.DB) *serialNumberRepository {
	return &serialNumberRepository{
		db: db,
	}
}

type serialNumberRepository struct {
	db *gorm.DB
}

func (s *serialNumberRepository) create(serialNumber *SerialNumber)  {
	 s.db.Create(serialNumber)
}
