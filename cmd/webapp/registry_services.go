package main

import (
	"template/pkg/di"
	"template/pkg/todo/delivery/http"

	"github.com/gin-gonic/gin"
)

func registryHTTPHandler(engin *gin.Engine) {
	toDoUseCases := di.NewToDoUseCases()
	toDoHandler := http.NewHandler(toDoUseCases)
	http.SetRoute(engin, toDoHandler)
}
