package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"projects/repository"
	"projects/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_ctrl_Transfer(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{}) // open gorm db
	assert.NoError(t, err)

	var columns []string
	columns = append(columns, []string{
		"account_number", "balance",
	}...)

	repo := repository.NewRepo(gdb)
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
		req, _ := http.NewRequest(http.MethodPost, cases.url, bytes.NewBuffer([]byte(cases.content)))
		req.Header.Set("Content-Type", "application/json")

		t.Run(cases.name, func(t *testing.T) {
			ctrl.Transfer(httptest.NewRecorder(), req)
		})
	}
}
