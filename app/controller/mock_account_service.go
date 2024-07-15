package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) GetAllAccount(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) GetAccountById(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) AddAccountData(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) UpdateAccountData(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) DeleteAccount(c *gin.Context) {
	m.Called(c)
}
