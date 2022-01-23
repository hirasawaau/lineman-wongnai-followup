package main

import (
	"lineman-wongnai-covid/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/covid/summary", handler.CovidSummary)

	app.Run()
}
