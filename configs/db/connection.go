package db

import (
	"log"

	// "github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=kadirbeksharau password=kadr2001 dbname=kadirbeksharau port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(
		&models.EntityUsers{},
		&models.EntityOrganizations{},
		&models.EntityCities{},
		&models.EntityFields{},
		&models.EntityFieldRates{},
		&models.EntityModerators{},
		&models.EntitySportType{},
		&models.EntityBookings{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}
 
	return db
}

func CloseDB(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.Close()
}