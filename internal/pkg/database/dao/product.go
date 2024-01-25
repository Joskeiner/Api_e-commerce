package dao

type Product struct {
	Model
	Name          string         `json:"ame"`
	Slug          string         `json:"slug"`
	ResellerPrice int            `json:"reseller_price"`
	RetailPrice   int            `json:"retail_price"`
	Stock         int            `json:"stock" `
	Description   string         `json:"descriptio" gorm:"type:text"`
	ShopID        uint           `json:"shop_id" gorm:"not null"`
	Shop          *Shop          `json:"shop"`
	CategoryID    uint           `json:"category_id" gorm:"not null"`
	Category      *Category      `json:"category"`
	Phontos       []ProductPhoto `json:"phontos" gorm:"foreingkey:ProductID"`
}

type ProductLog struct {
	Model
	ProductID     uint           `json:"Product_id" gorm:"not null"`
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	ResellerPrice int            `json:"reseller_price"`
	RetailPrice   int            `json:"retail_price"`
	Description   string         `json:"descriptio" gorm:"type:text"`
	ShopID        uint           `json:"shop_id" gorm:"not null"`
	Shop          *Shop          `json:"shop"`
	CategoryID    uint           `json:"category_id" gorm:"not null"`
	Category      *Category      `json:"category"`
	Phontos       []ProductPhoto `json:"phontos" gorm:"foreingkey:ProductID"`
}
type ProductPhoto struct {
	Model
	Url          string `json:"url"`
	ProductID    uint   `json:"Product_id" gorm:"not null;index"`
	ProductLogID uint   `json:"Product_log_id" gorm:"not null;index"`
}
