package main

import (
	"flag"
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	portNumber := flag.String("port", "8080", "Enter a port number")
	flag.Parse()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		fmt.Println("Received request from ", c.Request.URL.RequestURI())
		fmt.Printf("%s %s %s\n", c.Request.Method, c.Request.RequestURI, c.Request.Proto)
		fmt.Println("Host: ", c.Request.URL.Hostname())
		fmt.Println("User Agent: ", c.Request.UserAgent())
		fmt.Println("Accept: ", c.GetHeader("Accept"))
		c.String(200, "Hello From Backend Server")
		fmt.Println("Replied with a hello message")
	})

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(200, "Healthy")
	})
	router.Run(fmt.Sprintf(":%s", *portNumber)) // listen and serve on 0.0.0.0:8080
}

