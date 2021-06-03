package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"projects/adapter"

	"projects/repository"
	"projects/usecase"
)

func Test_ctrl_Transfer(t *testing.T) {
	db := adapter.DBSQL()

	repo := repository.NewRepo(db)
	usecase := usecase.NewUC(repo)
	ctrl := NewCtrl(usecase)

	s := []struct {
		name    string
		wantErr bool
		url     string
		content string
	}{
		{
			name:    "Success Case",
			wantErr: false,
			url:     "http://localhost:8080/account/555001/transfer",
			content: `{"to_account_number":"555002","amount":100}`,
		},
	}

	for _, cases := range s {
		req, _ := http.NewRequest(http.MethodPost, cases.url, nil)
		req.Header.Set("Content-Type", "application/json")

		t.Run(cases.name, func(t *testing.T) {
			ctrl.CheckSaldo(httptest.NewRecorder(), req)
		})
	}
}
