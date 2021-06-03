package usecase

import (
	"context"
	"net/http"
	"testing"

	"projects/adapter"
	"projects/model"
	"projects/repository"
)

func Test_uc_Transfer(t *testing.T) {
	db := adapter.DBSQL()

	repo := repository.NewRepo(db)
	usecase := NewUC(repo)
	type args struct {
		ctx      context.Context
		transfer model.Transfer
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
				ctx: context.TODO(),
				transfer: model.Transfer{
					FromAccounNumber: "555001",
					ToAccounNumber:   "555002",
					Amount:           100,
				},
			},
			want2:   "Succesfully transfer",
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
			want2:   "not enough saldo",
			want3:   http.StatusBadGateway,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
