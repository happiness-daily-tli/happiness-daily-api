package entity

import "gorm.io/gorm"

type Contents struct {
	gorm.Model
	Contents string
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Contents{})
}
