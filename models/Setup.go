package models

import (
	"fmt"

	"github.com/faisalfjri/gofiber-boilerplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := "root:123456@tcp(127.0.0.1:3306)/go_restapi_fiber"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config("DB_USERNAME"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	DB = db
}
