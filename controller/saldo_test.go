package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"projects/repository"
	"projects/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	env := map[string]string{
		"SERVER_PORT": "8080",
		"TIMEOUT":     "20",
		"DB_DRIVER":   "mysql",
		"DB_URI":      "root:@tcp(127.0.0.1:3306)/link-aja?charset=utf8&parseTime=True",
		"DB_USER":     "root",
		"DB_HOST":     "127.0.0.1",
		"DB_PASSWORD": "app123",
		"DB_NAME":     "link-aja",
	}

	for key, value := range env {
		os.Setenv(key, value)
	}

	// config.LoadConfig("link-aja-api")
}

func Test_ctrl_CheckSaldo(t *testing.T) {
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
	}{
		{
			name:    "Success Case",
			wantErr: false,
			url:     "http://localhost:8080/account/555001",
		},
		{
			name:    "error Case",
			wantErr: true,
			url:     "http://localhost:8080/account/555009",
		},
	}

	for _, cases := range s {
		req, _ := http.NewRequest(http.MethodGet, cases.url, nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		t.Run(cases.name, func(t *testing.T) {
			ctrl.CheckSaldo(httptest.NewRecorder(), req)
		})
	}
}
