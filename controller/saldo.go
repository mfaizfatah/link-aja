package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/link-aja/utils"
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
