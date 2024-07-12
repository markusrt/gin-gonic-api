package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUser(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) GetUserById(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) AddUserData(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) UpdateUserData(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) DeleteUser(c *gin.Context) {
	m.Called(c)
}
