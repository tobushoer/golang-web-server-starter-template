package http

import "github.com/gin-gonic/gin"

// SetRoute set route
func SetRoute(engin *gin.Engine, handler *ToDoHandler) {
	engin.POST("/api/v1/todo", handler.createToDoEndpoint)
}
