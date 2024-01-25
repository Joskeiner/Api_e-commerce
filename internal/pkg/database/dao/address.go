package dao

type Address struct {
	Model
	Title       string `json:"title"`
	Receiver    string `json:"receiver"`
	PhoneNumber string `json:"phone_number"`
	Datails     string `json:"details" gorm:"type:text"`
	UserID      uint   `json:"user_id" gorm:"not null"`
}
