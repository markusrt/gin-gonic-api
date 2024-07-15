package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountService interface {
	GetAll(c *gin.Context)
	Retrieve(c *gin.Context)
	CreateAccount(c *gin.Context)
	UpdateAccountData(c *gin.Context)
	Delete(c *gin.Context)
}

type LocalAccountServiceImpl struct {
	accountRepository repository.AccountRepository
}

func (u LocalAccountServiceImpl) UpdateAccountData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update account data by id")
	accountID, _ := strconv.Atoi(c.Param("accountID"))

	var request dao.Account
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.accountRepository.FindAccountById(accountID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Email = request.Email
	data.Name = request.Password
	data.Status = request.Status
	u.accountRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u LocalAccountServiceImpl) Retrieve(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get account by id")
	accountID, _ := strconv.Atoi(c.Param("accountID"))

	data, err := u.accountRepository.FindAccountById(accountID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u LocalAccountServiceImpl) CreateAccount(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data account")
	var request dao.Account
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data, err := u.accountRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u LocalAccountServiceImpl) GetAll(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data account")

	data, err := u.accountRepository.FindAllAccount()
	if err != nil {
		log.Error("Happened Error when find all account data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u LocalAccountServiceImpl) Delete(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data account by id")
	accountID, _ := strconv.Atoi(c.Param("accountID"))

	err := u.accountRepository.DeleteById(accountID)
	if err != nil {
		log.Error("Happened Error when try delete data account from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func AccountServiceInit(accountRepository repository.AccountRepository) *LocalAccountServiceImpl {
	return &LocalAccountServiceImpl{
		accountRepository: accountRepository,
	}
}
