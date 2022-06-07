package db

import (
	"log"

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
		&models.EntitySportTypes{},
		&models.EntityFields{},
		&models.EntityFieldRates{},
		&models.EntityModerators{},
		&models.EntityBookings{},
	)

	AccountsDataMigrator(db)
	CityDataMigrator("Алматы", db)
	CityDataMigrator("Нур-Султан", db)
	CityDataMigrator("Актау", db)
	CityDataMigrator("Атырау", db)
	CityDataMigrator("Жанаозен", db)
	CityDataMigrator("Павлодар", db)
	CityDataMigrator("Петропавловск", db)
	CityDataMigrator("Семей", db)
	CityDataMigrator("Туркестан", db)
	CityDataMigrator("Шымкент", db)
	CityDataMigrator("Талдыкорган", db)
	CityDataMigrator("Тараз", db)
	CityDataMigrator("Уральск", db)
	CityDataMigrator("Усть-Каменогорск", db)

	SportTypeDataMigrator("Football", db)
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