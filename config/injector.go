// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var accountServiceSet = wire.NewSet(service.AccountServiceInit,
	wire.Bind(new(service.AccountService), new(*service.LocalAccountServiceImpl)),
)

var accountRepoSet = wire.NewSet(repository.AccountRepositoryInit,
	wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryImpl)),
)

var accountCtrlSet = wire.NewSet(controller.AccountControllerInit,
	wire.Bind(new(controller.AccountController), new(*controller.AccountControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, accountCtrlSet, accountServiceSet, accountRepoSet)
	return nil
}
