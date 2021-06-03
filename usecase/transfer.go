package usecase

import (
	"context"
	"errors"

	"net/http"

	"projects/model"
)

func (r *uc) Transfer(ctx context.Context, transfer model.Transfer) (context.Context, interface{}, string, int, error) {
	var (
		res  interface{}
		msg  string
		code = http.StatusCreated
	)

	if transfer.FromAccounNumber == transfer.ToAccounNumber {
		return ctx, nil, "account number cannot same", http.StatusForbidden, errors.New("transaction_invalid")
	}

	saldo, err := r.query.GetSaldo(transfer.FromAccounNumber)
	if err != nil {
		return ctx, nil, "akun tidak ditemukan", http.StatusNotFound, err
	}

	if saldo.Balance < transfer.Amount {
		return ctx, nil, "not enough saldo", http.StatusBadRequest, errors.New("transaction_invalid")
	}

	err = r.query.InsertTransfer(transfer)
	if err != nil {
		return ctx, nil, "Terjadi kesalahan pada sisi server. Coba beberapa saat lagi, Terima Kasih!", http.StatusInternalServerError, err
	}

	res = "Succesfully transfer"
	return ctx, res, msg, code, nil
}
