package db

import (
	"log"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video models.Field)
	Update(video models.Field)
	Delete(video models.Field)
	FindAll() []models.Field
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
 
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	database, err := db.connection.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.Close()
}

func (db *database) Save(video models.Field) {
	db.connection.Create(&video)
}

func (db *database) Update(video models.Field) {
	db.connection.Save(&video)
}

func (db *database) Delete(video models.Field) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []models.Field {
	var videos []models.Field
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
