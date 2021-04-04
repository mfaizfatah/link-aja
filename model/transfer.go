package model

type Transfer struct {
	ID               int     `json:"id" gorm:"column:id"`
	FromAccounNumber string  `json:"from_account_number" gorm:"column:from_account_number"`
	ToAccounNumber   string  `json:"to_account_number" gorm:"column:to_account_number"`
	Amount           float32 `json:"amount" gorm:"column:amount"`
}
