package main

import (
	"os"

	"projects/adapter"
	"projects/config"
	"projects/controller"
	"projects/repository"
	"projects/router"
	"projects/usecase"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	service := "link-aja-api"

	config.LoadConfig(service)
}

func main() {
	db := adapter.DBSQL()

	repo := repository.NewRepo(db)
	uc := usecase.NewUC(repo)
	ctrl := controller.NewCtrl(uc)

	router := router.NewRouter(ctrl)
	router.Router(os.Getenv("SERVER_PORT"))
}
