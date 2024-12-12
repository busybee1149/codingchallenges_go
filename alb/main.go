package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	backends := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	var serverSelected = 0
	router.GET("/", func(contextCopy *gin.Context) {
		responseChannel := make(chan string)
		fmt.Println("Received request from ", contextCopy.Request.URL.RequestURI())
		fmt.Printf("%s %s %s\n", contextCopy.Request.Method, contextCopy.Request.RequestURI, contextCopy.Request.Proto)
		fmt.Println("Host: ", contextCopy.Request.URL.Hostname())
		fmt.Println("User Agent: ", contextCopy.Request.UserAgent())
		fmt.Println("Accept: ", contextCopy.GetHeader("Accept"))
		fmt.Println("Calling backend ", serverSelected)
		go callBackend(backends[serverSelected], responseChannel)
		contextCopy.String(200, <- responseChannel)
		defer close(responseChannel)
		serverSelected = (serverSelected + 1) % len(backends)
	})
	router.Run(":80") // listen and serve on 0.0.0.0:80
}

func callBackend(backendUrl string, responseChannel chan string) {
	backendClient := &http.Client{}

	response, err := backendClient.Get(backendUrl)
	if err != nil {
		fmt.Println("Issue while calling Backend")
	}
	defer response.Body.Close()
	
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in parsing backend response")
	}
	responseChannel <- string(responseBody)
}