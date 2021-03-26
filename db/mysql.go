package db

import (
	"example.com/medium/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var MySQL *gorm.DB

func ConnectMySQL() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("MYSQL: Couldn't connect to Database")
	}

	database.AutoMigrate(models.User{})

	MySQL = database
}
