package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"projects/model"
	"projects/utils"

	"github.com/go-chi/chi"
)

func (c *ctrl) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transfer := model.Transfer{}

	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		log.Print(err)
		utils.Response(ctx, w, false, http.StatusBadRequest, "badRequest")
		return
	}

	accNumber := chi.URLParam(r, "accNumber")

	transfer.FromAccounNumber = accNumber
	ctx, res, msg, st, err := c.uc.Transfer(ctx, transfer)
	if err != nil || st >= http.StatusBadRequest {
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}
