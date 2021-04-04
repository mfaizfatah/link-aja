package config

import (
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				service: "link-aja-api",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadConfig(tt.args.service)
		})
	}
}
