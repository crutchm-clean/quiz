package main

import (
	"Kahoot"
	"Kahoot/pkg/handler"
	"Kahoot/pkg/repository"
	"Kahoot/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("failed to init config")
	}
	db, err := repository.NewPostgresDb()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	hands := handler.NewHandler(services)
	srv := new(Kahoot.Server)
	if err := srv.Run(viper.GetString("port"), hands.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
