package database

import (
	"github.com/subham043/go-fiber/app/models"
	"github.com/subham043/go-fiber/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// Build Mysql connection URL.
	dbConnURL, _ := utils.ConnectionURLBuilder("mysql")

	db, err := gorm.Open(mysql.Open(dbConnURL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// // Migrate the schema
	db.AutoMigrate(&models.User{})

	DB = db
}
