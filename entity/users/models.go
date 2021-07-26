package entity

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Contents{})
}

// Migrate the schema of database if needed
//func AutoMigrate() {
//	db := common.GetDB()
//
//	db.AutoMigrate(&UserModel{})
//	db.AutoMigrate(&FollowModel{})
//}
