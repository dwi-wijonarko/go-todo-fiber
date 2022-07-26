package database

import (
	"fmt"
	"strconv"

	"github.com/dwi-wijonarko/go-fiber-todo/config"
	models "github.com/dwi-wijonarko/go-fiber-todo/src/models/todo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Printf("Error parsing DB_PORT: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	//migrate the schema
	DB.AutoMigrate(&models.Todo{})
	fmt.Println("Database Migrated")

}
