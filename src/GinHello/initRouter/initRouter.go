package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"

)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("/Users/allanyang/gofile/src/GinHello/templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/Users/allanyang/gofile/src/GinHello/statics", "/Users/allanyang/gofile/src/GinHello/./statics/")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}
	userRouter := router.Group("/user")
	{
		//userRouter.GET("/:name",handler.UserSave)
		//userRouter.GET("",handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}
	return router
}

