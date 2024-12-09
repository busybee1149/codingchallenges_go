package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		fmt.Println("Received request from ", c.Request.URL.RequestURI())
		fmt.Printf("%s %s %s\n", c.Request.Method, c.Request.RequestURI, c.Request.Proto)
		fmt.Println("Host: ", c.Request.URL.Hostname())
		fmt.Println("User Agent: ", c.Request.UserAgent())
		fmt.Println("Accept: ", c.GetHeader("Accept"))

	})
	router.Run(":80") // listen and serve on 0.0.0.0:80
}