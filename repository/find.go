package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mfaizfatah/link-aja/model"
)

func (u *repo) GetSaldo(accNumber string) (*model.Saldo, error) {
	var saldo model.Saldo

	query := fmt.Sprintf(`SELECT account.account_number, customer.name, account.balance FROM account INNER JOIN customer ON account.customer_number = customer.customer_number WHERE account.account_number = %v`, accNumber)
	err := u.db.Raw(query).Scan(&saldo).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data_not_found")
	}

	return &saldo, err
}
