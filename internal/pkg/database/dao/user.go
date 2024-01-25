package dao

import (
	"time"
)

type User struct {
	Model
	Name         string        `json:"name"`
	Password     string        `json:"password"`
	PhoneNumber  string        `json:"phone_number" gorm:"unique;not null"`
	Email        string        `json:"email" gorm:"unique;not null"`
	BirthDate    time.Time     `json:"bith_date"`
	About        string        `json:"about"`
	Job          string        `json:"job"`
	CityID       string        `json:"city_id"`
	Addresses    []Address     `json:"address"`
	Shop         *Shop         `gorm:"foreignkey:UserID"`
	IsAdmin      bool          `json:"is_admin" gorm:"default:false"`
	Transactions []Transaction `json:"transaction"`
}
