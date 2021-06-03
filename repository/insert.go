package repository

import (
	"projects/model"
)

func (r *repo) InsertTransfer(transfer model.Transfer) error {
	err := r.db.Table(TableTransfer).Create(&transfer).Error
	if err != nil {
		return err
	}

	return nil
}
