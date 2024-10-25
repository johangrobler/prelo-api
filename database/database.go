package database

import (
	"fmt"
	"log"
	"os"
	"prelo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=prefer TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	dsn := config.Config("DB")

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to pg")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = Dbinstance{
		Db: db,
	}
}
