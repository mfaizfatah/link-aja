package model

type Saldo struct {
	AccounNumber string  `json:"account_number" gorm:"column:account_number"`
	Balance      float32 `json:"balance" gorm:"column:balance"`
	Name         string  `json:"customer_name" gorm:"column:name"`
}
