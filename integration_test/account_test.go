package integration_test

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUsers_ReturnsList(t *testing.T) {
	init := config.Init()
	app := router.Init(init)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/users", nil)
	app.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
}
