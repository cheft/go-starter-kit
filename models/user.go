package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	db *gorm.DB
	gorm.Model
	Code  string
	Price uint
}

func NewProduct() Product {
	db := DB()
	// Migrate the schema
	db.AutoMigrate(&Product{})

	model := Product{db: db}
	return model
}

func (p Product) GetAll() *gorm.DB {
	return p.db.Find(&Product{})
}

func (p Product) Create() {
	p.db.Create(&Product{Code: "L1212", Price: 1000})
}
