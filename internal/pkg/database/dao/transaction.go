package dao

type Transaction struct {
	Model
	TotalPrice         uint                `json:"total_price"`
	Invoice            string              `json:"invoice" gorm:"unique;not null"`
	PaymentMethod      string              `json:"payment_method"`
	UserID             uint                `json:"user_id"`
	User               User                `json:"user"`
	AddressID          uint                `json:"address_id"`
	Address            Address             `json:"address"`
	TransactionDetails []TransactionDetail `json:"transaction_detail"`
}

type TransactionDetail struct {
	Model
	Quantity      uint        `json:"quantity"`
	TotalPrice    uint        `json:"total_price"`
	TransactionID uint        `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
	ProductLogID  uint        `json:"product_log_id"`
	ProductLog    ProductLog  `json:"product_log"`
	ShopID        uint        `json:"shop_id"`
	Shop          Shop        `json:"shop"`
}
