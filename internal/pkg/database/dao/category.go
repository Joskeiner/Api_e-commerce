package dao

type Category struct {
	Model
	Name     string    `json:"name" gorm:"unique;not null"`
	Products []Product `json:"product"`
}
