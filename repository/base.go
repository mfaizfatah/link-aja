package repository

import (
	"projects/model"

	"gorm.io/gorm"
)

const (
	// TableTransfer is name for customer table on db
	TableTransfer = "transfer"
)

// repo struct with value mysqldb connection
type repo struct {
	db *gorm.DB
}

// Repo represent the Repository contract
type Repo interface {
	//find
	GetSaldo(accNumber string) (*model.Saldo, error)
	//insert
	InsertTransfer(transfer model.Transfer) error
}

/*NewRepo will create an object that represent the Repository interface (Repo)
 * @parameter
 * db - mysql database connection
 *
 * @represent
 * interface Repo
 *
 * @return
 * repo struct with value db (mysql database connection)
 */
func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
