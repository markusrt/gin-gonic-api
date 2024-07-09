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
		users := api.Group("/users")
		{
			users.GET("", init.UserCtrl.GetAllUserData)
			users.POST("", init.UserCtrl.AddUserData)
			users.GET("/:userID", init.UserCtrl.GetUserById)
			users.PUT("/:userID", init.UserCtrl.UpdateUserData)
			users.DELETE("/:userID", init.UserCtrl.DeleteUser)
		}
	}

	return router
}
