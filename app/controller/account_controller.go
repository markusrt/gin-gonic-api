package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAllAccountData(c *gin.Context)
	AddAccountData(c *gin.Context)
	GetAccountById(c *gin.Context)
	UpdateAccountData(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type AccountControllerImpl struct {
	svc service.AccountService
}

func (u AccountControllerImpl) GetAllAccountData(c *gin.Context) {
	u.svc.GetAllAccount(c)
}

func (u AccountControllerImpl) AddAccountData(c *gin.Context) {
	u.svc.AddAccountData(c)
}

func (u AccountControllerImpl) GetAccountById(c *gin.Context) {
	u.svc.GetAccountById(c)
}

func (u AccountControllerImpl) UpdateAccountData(c *gin.Context) {
	u.svc.UpdateAccountData(c)
}

func (u AccountControllerImpl) DeleteAccount(c *gin.Context) {
	u.svc.DeleteAccount(c)
}

func AccountControllerInit(accountService service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{
		svc: accountService,
	}
}
