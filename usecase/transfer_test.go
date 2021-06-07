package usecase

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net/http"
	"testing"

	"projects/model"
	"projects/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_uc_Transfer(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
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
	usecase := NewUC(repo)

	type args struct {
		ctx      context.Context
		transfer model.Transfer
	}
	tests := []struct {
		name    string
		args    args
		rows    driver.Result
		want2   string
		want3   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				transfer: model.Transfer{
					FromAccounNumber: "555001",
					ToAccounNumber:   "555002",
					Amount:           100,
				},
			},
			rows:    sqlmock.NewResult(1, 1),
			want2:   "",
			want3:   http.StatusCreated,
			wantErr: false,
		},
		{
			name: "same account number from and to",
			args: args{
				ctx: context.TODO(),
				transfer: model.Transfer{
					FromAccounNumber: "555001",
					ToAccounNumber:   "555001",
					Amount:           100,
				},
			},
			rows:    sqlmock.NewResult(1, 1),
			want2:   "account number cannot same",
			want3:   http.StatusForbidden,
			wantErr: true,
		},
		{
			name: "not enough saldo",
			args: args{
				ctx: context.TODO(),
				transfer: model.Transfer{
					FromAccounNumber: "555001",
					ToAccounNumber:   "555002",
					Amount:           1000000,
				},
			},
			rows:    sqlmock.NewResult(1, 1),
			want2:   "not enough saldo",
			want3:   http.StatusBadRequest,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := fmt.Sprintf(`SELECT account.account_number, customer.name, account.balance FROM account INNER JOIN customer ON account.customer_number = customer.customer_number WHERE account.account_number = %v`, 555001)
			mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows(columns).AddRow("555001", 10000))

			query = "INSERT INTO `transfer` (`from_account_number`,`to_account_number`,`amount`) VALUES (?,?,?)"
			mock.ExpectExec(query).
				WithArgs(tt.args.transfer.FromAccounNumber, tt.args.transfer.ToAccounNumber, tt.args.transfer.Amount).
				WillReturnResult(tt.rows)

			_, _, got2, got3, err := usecase.Transfer(tt.args.ctx, tt.args.transfer)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got2 != tt.want2 {
				t.Errorf("uc.Transfer() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("uc.Transfer() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
