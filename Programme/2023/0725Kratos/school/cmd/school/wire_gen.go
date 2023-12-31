// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"school/internal/biz"
	"school/internal/conf"
	"school/internal/data"
	"school/internal/server"
	"school/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewGormDB(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	taskRepo := data.NewTaskRepo(dataData, logger)
	taskUsecase := biz.NewTaskUsecase(taskRepo, logger)
	schoolService := service.NewSchoolService(taskUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, schoolService, logger)
	httpServer := server.NewHTTPServer(confServer, schoolService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
