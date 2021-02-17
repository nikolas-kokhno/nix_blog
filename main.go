package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	/* Initialize config */
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to load config file: %s", err.Error())
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
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}
