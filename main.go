package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomdsg/actions-with-docker-example/healthcheck"
)

func main() {
	e := gin.Default()
	e.GET("/healthcheck", healthcheck.NewHandlerFn())
	e.Run(":3000")
}
