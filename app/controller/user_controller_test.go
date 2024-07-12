package controller

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Context, *httptest.ResponseRecorder, *MockUserService) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	request, _ := http.NewRequest("GET", "/", nil)
	context.Request = request
	mockUserService := new(MockUserService)
	return context, recorder, mockUserService
}

func Test_AddUserData_CallsUserService(t *testing.T) {
	context, recorder, mockUserService := setupTest()
	mockUserService.On("AddUserData", context).Return()
	userControllerImpl := UserControllerInit(mockUserService)

	userControllerImpl.AddUserData(context)

	mockUserService.AssertCalled(t, "AddUserData", context)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_AddUserData_DeleteUser(t *testing.T) {
	context, recorder, mockUserService := setupTest()
	mockUserService.On("DeleteUser", context).Return()
	userControllerImpl := UserControllerInit(mockUserService)

	userControllerImpl.DeleteUser(context)

	mockUserService.AssertCalled(t, "DeleteUser", context)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	return c, w
}
