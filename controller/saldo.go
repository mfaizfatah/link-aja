package controller

import (
	"net/http"

	"projects/utils"

	"github.com/go-chi/chi"
)

func (c *ctrl) CheckSaldo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accNumber := chi.URLParam(r, "accNumber")

	ctx, res, msg, st, err := c.uc.CheckSaldo(ctx, accNumber)
	if err != nil || st >= http.StatusBadRequest {
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}
