package usecase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"projects/repository"

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

func Test_uc_CheckSaldo(t *testing.T) {
	db, mock, err := sqlmock.New() // mock sql.DB
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

	repo := repository.NewRepo(gdb)

	var columns []string
	columns = append(columns, []string{
		"account_number", "balance",
	}...)

	usecase := NewUC(repo)

	query := fmt.Sprintf(`SELECT account.account_number, customer.name, account.balance FROM account INNER JOIN customer ON account.customer_number = customer.customer_number WHERE account.account_number = %v`, 555001)
	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows(columns).AddRow("555001", 10000))

	type args struct {
		ctx       context.Context
		accNumber string
	}
	tests := []struct {
		name    string
		args    args
		want2   string
		want3   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx:       context.TODO(),
				accNumber: "555001",
			},
			want2:   "",
			want3:   http.StatusOK,
			wantErr: false,
		},
		{
			name: "not_found",
			args: args{
				ctx:       context.TODO(),
				accNumber: "555009",
			},
			want2:   "akun tidak ditemukan",
			want3:   http.StatusNotFound,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, got2, got3, err := usecase.CheckSaldo(tt.args.ctx, tt.args.accNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CheckSaldo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got2 != tt.want2 {
				t.Errorf("uc.CheckSaldo() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("uc.CheckSaldo() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
