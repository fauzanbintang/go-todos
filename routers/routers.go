package routers

import (
	"golang-ptp/go-todos/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todos := router.Group("/todos")
	{
		todos.GET("/", controllers.GetTodos)
		todos.GET("/:id", controllers.GetTodo)
		todos.POST("/", controllers.CreateTodo)
		todos.PUT("/:id", controllers.PutTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
