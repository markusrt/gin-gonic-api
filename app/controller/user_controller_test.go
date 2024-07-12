package controller

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func AddUserData_CallsUserService(t *testing.T) {
	context, w := GetTestGinContext()
	mockUserService := new(MockUserService)
	mockUserService.On("AddUserData", context).Return()
	userControllerImpl := UserControllerInit(mockUserService)

	userControllerImpl.AddUserData(context)

	mockUserService.AssertCalled(t, "AddUserData", context)
	assert.Equal(t, http.StatusOK, w.Code)
}

func AddUserData_DeleteUser(t *testing.T) {
	context, w := GetTestGinContext()
	mockUserService := new(MockUserService)
	mockUserService.On("DeleteUser", context).Return()
	userControllerImpl := UserControllerInit(mockUserService)

	userControllerImpl.DeleteUser(context)

	mockUserService.AssertCalled(t, "DeleteUser", context)
	assert.Equal(t, http.StatusOK, w.Code)
}

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	return c, w
}
