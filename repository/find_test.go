package repository

import (
	"os"
	"reflect"
	"testing"

	"projects/adapter"
	"projects/config"
	"projects/model"
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

	config.LoadConfig("link-aja-api")
}

func Test_repo_GetSaldo(t *testing.T) {
	db := adapter.DBSQL()

	repo := NewRepo(db)

	type args struct {
		accNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Saldo
		wantErr bool
	}{
		// TODO: Add test cases
		{
			name: "found",
			args: args{
				accNumber: "555001",
			},
			want: &model.Saldo{
				AccounNumber: "555001",
				Balance:      10000,
				Name:         "Bob Martin",
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				accNumber: "555009",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetSaldo(tt.args.accNumber)
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
