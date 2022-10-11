package repository

import (
	"gintoki/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entity.Video) error
	Update(video entity.Video) error
	Delete(video entity.Video) error
	FindAll() ([]entity.Video, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) error {
	return db.connection.Create(&video).Error
}

func (db *database) Update(video entity.Video) error {
	return db.connection.Model(&video).Update(&video).Error
}

func (db *database) Delete(video entity.Video) error {
	return db.connection.Delete(&video).Error
}

func (db *database) FindAll() ([]entity.Video, error) {
	var videos []entity.Video

	return videos, db.connection.Set("gorm:auto_preload", true).Find(&videos).Error
}
