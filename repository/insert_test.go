package repository

import (
	"testing"

	"github.com/mfaizfatah/link-aja/adapter"
	"github.com/mfaizfatah/link-aja/model"
)

func Test_repo_InsertTransfer(t *testing.T) {
	db := adapter.DBSQL()

	repo := NewRepo(db)
	type args struct {
		transfer model.Transfer
	}
	tests := []struct {
		name    string
		args    args
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.InsertTransfer(tt.args.transfer); (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertTransfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
