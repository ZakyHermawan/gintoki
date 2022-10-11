package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id" example:"1"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)" example:"zaky"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)" example:"hermawan"`
	Age       int    `json:"age" binding:"gte=1,lte=130" example:"20"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)" example:"zaky.hermawan9615@gmail.com"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id" example:"1"`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)" example:"Hello World"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)" example:"Introduction"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);" example:"google.com"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
