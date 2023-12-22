package main

import (
	"github.com/D-building-anonymaizer/backend-service"
	"github.com/D-building-anonymaizer/backend-service/pkg/handler"
	"github.com/D-building-anonymaizer/backend-service/pkg/repository"
	"github.com/D-building-anonymaizer/backend-service/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while initializing the config file: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(backend.Server)
	if err := srv.Run(viper.GetString("port"), viper.GetString("ip"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
