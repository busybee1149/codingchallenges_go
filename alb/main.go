package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(contextCopy *gin.Context) {
		responseChannel := make(chan string)
		fmt.Println("Received request from ", contextCopy.Request.URL.RequestURI())
		fmt.Printf("%s %s %s\n", contextCopy.Request.Method, contextCopy.Request.RequestURI, contextCopy.Request.Proto)
		fmt.Println("Host: ", contextCopy.Request.URL.Hostname())
		fmt.Println("User Agent: ", contextCopy.Request.UserAgent())
		fmt.Println("Accept: ", contextCopy.GetHeader("Accept"))
		go callBackend(responseChannel)
		contextCopy.String(200, <- responseChannel)
		defer close(responseChannel)
	})
	router.Run(":80") // listen and serve on 0.0.0.0:80
}

func callBackend(responseChannel chan string) {
	backendClient := &http.Client{}

	response, err := backendClient.Get("http://localhost:8080")
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