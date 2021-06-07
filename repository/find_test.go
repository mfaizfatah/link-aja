package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"projects/model"

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

func Test_repo_GetSaldo(t *testing.T) {
	// db := adapter.DBSQL()

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

	repo := NewRepo(gdb)

	var columns []string
	columns = append(columns, []string{
		"account_number", "balance",
	}...)

	type args struct {
		accNumber string
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    *model.Saldo
		wantErr bool
	}{
		// TODO: Add test cases
		{
			name: "found",
			args: args{
				accNumber: "555001",
			},
			rows: sqlmock.NewRows(columns).AddRow("555001", 10000),
			want: &model.Saldo{
				AccounNumber: "555001",
				Balance:      10000,
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				accNumber: "9999",
			},
			rows: sqlmock.NewRows(columns).AddRow("5", "one").RowError(2, errors.New("nil documents")),
			want: &model.Saldo{
				AccounNumber: "5",
				Balance:      0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			query := fmt.Sprintf(`SELECT account.account_number, customer.name, account.balance FROM account INNER JOIN customer ON account.customer_number = customer.customer_number WHERE account.account_number = %v`, tt.args.accNumber)
			mock.ExpectQuery(query).WillReturnRows(tt.rows)

			got, err := repo.GetSaldo(tt.args.accNumber)
			log.Print(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetSaldo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetSaldo() = %v, want %v", got, tt.want)
			}

		})
	}
}
