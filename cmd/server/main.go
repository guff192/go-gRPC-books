package main

import (
	"github.com/guff192/go-gRPC-books/pkg/repository"
	"github.com/guff192/go-gRPC-books/pkg/server"
	"github.com/guff192/go-gRPC-books/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	// initialize config for DB
	db, err := repository.NewMysqlDB(repository.Config{
		Host:     "localhost",
		Port:     "3308",
		Username: "root",
		Password: "qwerty",
		DBName:   "library",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// creating repository and service
	repos := repository.NewRepository(db)
	service := service.NewService(repos)

	// starting server
	if err = server.RunServer(server.Config{Host: "localhost", Port: "8080"}, service); err != nil {
		logrus.Fatalf("failed to serve: %s", err)
	}
}
