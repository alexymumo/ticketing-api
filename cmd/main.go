package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", homepage)
	router.Run()
}

func homepage(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}
