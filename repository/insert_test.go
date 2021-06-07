package repository

import (
	"database/sql/driver"
	"testing"

	"projects/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_repo_InsertTransfer(t *testing.T) {
	// db := adapter.DBSQL()

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

	repo := NewRepo(gdb)

	var columns []string
	columns = append(columns, []string{
		"account_number", "balance",
	}...)

	type args struct {
		transfer model.Transfer
	}
	tests := []struct {
		name    string
		args    args
		rows    driver.Result
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				transfer: model.Transfer{
					FromAccounNumber: "555001",
					ToAccounNumber:   "555002",
					Amount:           100,
				},
			},
			rows: sqlmock.NewResult(1, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "INSERT INTO `transfer` (`from_account_number`,`to_account_number`,`amount`) VALUES (?,?,?)"
			mock.ExpectExec(query).
				WithArgs(tt.args.transfer.FromAccounNumber, tt.args.transfer.ToAccounNumber, tt.args.transfer.Amount).
				WillReturnResult(tt.rows)

			if err := repo.InsertTransfer(tt.args.transfer); (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertTransfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
