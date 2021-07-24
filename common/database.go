package common

import (
	"fmt"
	_ "fmt"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// gorm를 사용하여 DB를 Open 한다.
func Init() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "./../gorm.db")
	//if err != nil {
	//	fmt.Println("db err: (Init) ", err)
	//}
	//db.DB().SetMaxIdleConns(10)
	////db.LogMode(true)
	//DB = db
	//return DB
}
