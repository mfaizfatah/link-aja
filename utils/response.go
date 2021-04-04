package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// response using for sending response to frontend
type response struct {
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// Response is sending data response to frontend
func Response(ctx context.Context, w http.ResponseWriter, status bool, statuscode int, data interface{}) {
	var (
		res response
	)

	if !status {
		res.Status = "error"
		res.ErrorMessage = data.(string)
		res.Data = ""
	} else {
		res.Status = "success"
		res.ErrorMessage = ""
		res.Data = data
	}

	datares, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	fmt.Fprintf(w, string(datares))
}

// HTMLResponse for response html
func HTMLResponse(w http.ResponseWriter, statuscode int, AddForm string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statuscode)
	fmt.Fprint(w, AddForm)
	return
}
