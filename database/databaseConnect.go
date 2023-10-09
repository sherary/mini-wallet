package database

import (
	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"

	"fmt"
	helpers "mini-wallet/helpers"
	models "mini-wallet/models"
)

var db *gorm.DB

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.WalletManager{}, &models.TransactionHistory{})
}

func PopulateTable(db *gorm.DB) {
	for i := 0; i <= 15; i++ {
		db.Create(&models.User{
			Name:         faker.Name(),
			Username:     faker.Username(),
			Password:     faker.Password(),
			Customer_XID: []byte(xid.New().String()),
		})
	}
}

func DatabaseConnect() {
	var err error
	// Retrieve environment variables for database connection
	dbUsername := helpers.GoDotENVLoader("DB_USERNAME")
	dbPassword := helpers.GoDotENVLoader("DB_PASSWORD")
	dbHost := helpers.GoDotENVLoader("DB_HOST")
	dbPort := helpers.GoDotENVLoader("DB_PORT")
	dbName := helpers.GoDotENVLoader("DB_NAME")

	// Create DSN string using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	CreateTable(db)

	var user models.User
	if db.First(&user).Error == gorm.ErrRecordNotFound {
		PopulateTable(db)
	}
}
