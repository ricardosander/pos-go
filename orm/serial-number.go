package main

type SerialNumber struct {
	ID int `gorm:"primaryKey"`
	Number string
	ProductID int
}
