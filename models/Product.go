package models

import (
	"time"
)

// Data Product
type Product struct {

	Id int `gorm:"PrimaryKey" json:"Id"`
	Name_book string `gorm:"type:varchar(25)" json:"Name_book"`
	Author string `gorm:"type:varchar(25)" json:"Author"`
	CreatedAt time.Time
	UpdatedAt time.Time


}