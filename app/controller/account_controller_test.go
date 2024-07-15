package controller

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Context, *httptest.ResponseRecorder, *MockAccountService) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	request, _ := http.NewRequest("GET", "/", nil)
	context.Request = request
	mockAccountService := new(MockAccountService)
	return context, recorder, mockAccountService
}

func Test_AddAccountData_CallsAccountService(t *testing.T) {
	context, recorder, mockAccountService := setupTest()
	mockAccountService.On("AddAccountData", context).Return()
	accountControllerImpl := AccountControllerInit(mockAccountService)

	accountControllerImpl.AddAccountData(context)

	mockAccountService.AssertCalled(t, "AddAccountData", context)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_AddAccountData_DeleteAccount(t *testing.T) {
	context, recorder, mockAccountService := setupTest()
	mockAccountService.On("DeleteAccount", context).Return()
	accountControllerImpl := AccountControllerInit(mockAccountService)

	accountControllerImpl.DeleteAccount(context)

	mockAccountService.AssertCalled(t, "DeleteAccount", context)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	return c, w
}
