// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
	"github.com/google/wire"
)

import (
	_ "github.com/ncruces/go-sqlite3/embed"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	accountRepositoryImpl := repository.AccountRepositoryInit(gormDB)
	LocalAccountServiceImpl := service.AccountServiceInit(accountRepositoryImpl)
	accountControllerImpl := controller.AccountControllerInit(LocalAccountServiceImpl)
	initialization := NewInitialization(accountRepositoryImpl, LocalAccountServiceImpl, accountControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var accountServiceSet = wire.NewSet(service.AccountServiceInit, wire.Bind(new(service.AccountService), new(*service.LocalAccountServiceImpl)))

var accountRepoSet = wire.NewSet(repository.AccountRepositoryInit, wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryImpl)))

var accountCtrlSet = wire.NewSet(controller.AccountControllerInit, wire.Bind(new(controller.AccountController), new(*controller.AccountControllerImpl)))
