package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mfaizfatah/link-aja/adapter"
	"github.com/mfaizfatah/link-aja/config"
	"github.com/mfaizfatah/link-aja/controller"
	"github.com/mfaizfatah/link-aja/repository"
	"github.com/mfaizfatah/link-aja/router"
	"github.com/mfaizfatah/link-aja/usecase"
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
