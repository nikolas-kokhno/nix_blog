package main

import (
	"log"
	"os"

	"github.com/nikolas-kokhno/nix_blog/models"
	"github.com/nikolas-kokhno/nix_blog/routers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// @title Blog API
// @version 1.0
// @description REST API server from NIX education

// @host localhost:8181
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	/* Initialize config */
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to load config file: %s", err.Error())
	}

	/* Initialize database and start migration */
	if len(os.Args) > 1 && os.Args[1] == "--initDB" {
		models.InitDatabase()
	}

	/* Create and start server */
	e := echo.New()
	Run(e)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func Run(e *echo.Echo) {
	routers.InitRoutes(e)
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}
