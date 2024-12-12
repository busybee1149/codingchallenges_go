package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


func main() {
	healthCheckIntervalInSeconds := flag.Int("healthCheckIntervalInSeconds", 10, "Specify a non-negative interval in seconds within (1-100) for health checks")
	flag.Parse()

	if (*healthCheckIntervalInSeconds < 1 && *healthCheckIntervalInSeconds > 100) {
		fmt.Println("Invalid Health Check Interval")
		os.Exit(-1)
	}
	
	tickerDuration := time.Second * time.Duration(*healthCheckIntervalInSeconds)
	ticker := time.NewTicker(tickerDuration)
	
	router := gin.Default()
	backends := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	go healthCheck(backends, ticker)
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

func healthCheck(servers []string, ticker *time.Ticker) {
	for range ticker.C {
		for _, server := range servers {
			healthCheckUrl := fmt.Sprintf("%s/healthcheck", server)
			backendClient := &http.Client{}
			response, err := backendClient.Get(healthCheckUrl)
			if err != nil {
				fmt.Println("Issue while calling Backend")
				continue
			}
			if response.StatusCode != 200 {
				fmt.Printf("%s is unhealthy", server)
			}
		}
	}
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