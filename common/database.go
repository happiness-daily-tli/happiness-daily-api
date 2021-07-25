package common

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Init gorm 를 사용하여 DB를 Open 한다.
func Init() *gorm.DB {

	fmt.Println("DB OPEN")
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db, err := gorm.Open("sqlite3", "./../gorm.db")
	//if err != nil {
	//	fmt.Println("db err: (Init) ", err)
	//}
	//db.DB().SetMaxIdleConns(10)
	////db.LogMode(true)
	//DB = db
	//return DB
}
