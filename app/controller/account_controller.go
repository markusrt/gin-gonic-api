package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	AccountSummary(c *gin.Context)
	CreateAccount(c *gin.Context)
	AccountDetails(c *gin.Context)
	UpdateAccountData(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type AccountControllerImpl struct {
	svc service.AccountService
}

func (u AccountControllerImpl) AccountSummary(c *gin.Context) {
	u.svc.GetAll(c)
}

func (u AccountControllerImpl) CreateAccount(c *gin.Context) {
	u.svc.CreateAccount(c)
}

func (u AccountControllerImpl) AccountDetails(c *gin.Context) {
	u.svc.Retrieve(c)
}

func (u AccountControllerImpl) UpdateAccountData(c *gin.Context) {
	u.svc.UpdateAccountData(c)
}

func (u AccountControllerImpl) DeleteAccount(c *gin.Context) {
	u.svc.Delete(c)
}

func AccountControllerInit(accountService service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{
		svc: accountService,
	}
}
