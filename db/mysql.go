package db

import (
	"fmt"
	"os"

	"example.com/medium/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MySQL *gorm.DB

func ConnectMySQL() {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")

	database, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/user", user, pass, host))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(models.User{})

	MySQL = database
}
