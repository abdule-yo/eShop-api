package database

import (
	"log"
	"os"

	"github.com/abdule-yo/eCommerce-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	//Todo: connection database
	dns := "postgresql://eCommerce-api-go_owner:npg_lOWrG2ULm0wv@ep-square-flower-a7854eed-pooler.ap-southeast-2.aws.neon.tech/eCommerce-api-go?sslmode=require"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	//Todo: checking errors
	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected successfully")

	//Todo: logging the info (anything happening inside the connection to the DB)s
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//Todo: Migrations (creating empty tables of each schema)
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	//Todo: exporting database for global usage
	Database = DbInstance{Db: db}

}
