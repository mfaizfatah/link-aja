package config

import (
	"os"
	"strings"

	"projects/adapter"

	_ "github.com/joho/godotenv/autoload"
)

/*LoadConfig is for load all configuration, connection etc
 * @parameters
 * service is for service name
 */
func LoadConfig(service string) {
	dbdriver := strings.ToLower(os.Getenv("DB_DRIVER"))
	if dbdriver == "mysql" {
		sql := os.Getenv("DB_URI")
		adapter.LoadMySQL(sql)
	}

	// utils.Newprometheus(service)
}
