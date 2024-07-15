package router

import (
	"gin-gonic-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		accounts := api.Group("/accounts")
		{
			accounts.GET("", init.AccountCtrl.GetAllAccountData)
			accounts.POST("", init.AccountCtrl.AddAccountData)
			accounts.GET("/:accountID", init.AccountCtrl.GetAccountById)
			accounts.PUT("/:accountID", init.AccountCtrl.UpdateAccountData)
			accounts.DELETE("/:accountID", init.AccountCtrl.DeleteAccount)
		}
	}

	return router
}
