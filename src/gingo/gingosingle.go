package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorldGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World in GET")
}

func helloWorlPost(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World in POST")
}

func fetchId(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("id is :%s\n", id))
}

func action1(c *gin.Context) {
	c.String(http.StatusOK, "action1")
}

func action2(c *gin.Context) {
	c.String(http.StatusOK, "action2")
}

func action3(c *gin.Context) {
	c.String(http.StatusOK, "action3")
}

func main() {
	router := gin.Default()

	router.GET("/helloWorld", helloWorldGet)

	router.POST("/helloWorld", helloWorlPost)

	router.GET("/param/:id", fetchId)

	//组路由
	group1 := router.Group("/g1")
	{
		group1.GET("action1", action1)
		group1.GET("action2", action2)
		group1.GET("action3", action3)
	}

	router.Run("127.0.0.1:8083")
}
