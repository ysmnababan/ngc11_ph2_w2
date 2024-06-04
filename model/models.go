package model

type User struct {
	UserID        uint    `json:"user_id" gorm:"primaryKey"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

type Product struct {
	ProductID uint    `gorm:"primaryKey" json:"product_id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	UserID        uint
	ProductID     uint
	Quantity      int
	TotalAmount   float64
}
