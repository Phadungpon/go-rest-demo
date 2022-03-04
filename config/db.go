package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDb() {

	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(gin.Mode() == gin.DebugMode)
}

func getDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
