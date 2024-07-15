package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) GetAll(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) Retrieve(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) CreateAccount(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) UpdateAccountData(c *gin.Context) {
	m.Called(c)
}

func (m *MockAccountService) Delete(c *gin.Context) {
	m.Called(c)
}
