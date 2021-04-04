package usecase

import (
	"context"
	"net/http"
)

func (r *uc) CheckSaldo(ctx context.Context, accNumber string) (context.Context, interface{}, string, int, error) {
	var (
		res  interface{}
		msg  string
		code = http.StatusOK
	)

	saldo, err := r.query.GetSaldo(accNumber)
	if err != nil {
		return ctx, nil, "akun tidak ditemukan", http.StatusNotFound, err
	}

	res = saldo

	return ctx, res, msg, code, nil
}
