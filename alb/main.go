package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"errors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	endpointUrl string
	isHealthy bool
}

type ServerResponse struct {
	statusCode int
	responseBody string
}

func NewServer(endpointUrl string) Server {
	return Server{
		endpointUrl: endpointUrl,
		isHealthy: true,
	}
}

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
	backends := []Server{
		NewServer("http://localhost:8080"),
		NewServer("http://localhost:8081"),
		NewServer("http://localhost:8082"),
	}
	go healthCheck(&backends, ticker)

	var serverToSelect = 0
	router.GET("/", func(contextCopy *gin.Context) {
		selectedServer, err := selectServer(serverToSelect, &backends)
		if err != nil {
			contextCopy.String(503, "Server unavailable")
		} else {
			responseChannel := make(chan ServerResponse)
			fmt.Println("Received request from ", contextCopy.Request.URL.RequestURI())
			fmt.Printf("%s %s %s\n", contextCopy.Request.Method, contextCopy.Request.RequestURI, contextCopy.Request.Proto)
			fmt.Println("Host: ", contextCopy.Request.URL.Hostname())
			fmt.Println("User Agent: ", contextCopy.Request.UserAgent())
			fmt.Println("Accept: ", contextCopy.GetHeader("Accept"))
			fmt.Println("Calling backend ", selectedServer.endpointUrl)
			go callBackend(selectedServer.endpointUrl, responseChannel)
			serverResponse := <- responseChannel
			contextCopy.String(serverResponse.statusCode, serverResponse.responseBody)
			defer close(responseChannel)
			serverToSelect = (serverToSelect + 1) % len(*&backends)
		}
	})
	router.Run(":80") // listen and serve on 0.0.0.0:80
}

func selectServer(serverToSelect int, backends *[]Server) (Server, error) {
	var currentIndex int
	totalServers := len(*backends)

	for index := range totalServers {
		currentIndex = (index + serverToSelect) % totalServers
		if (*backends)[currentIndex].isHealthy {
			return (*backends)[currentIndex], nil
		}
	}
	return Server{}, errors.New("no server available")
}

func healthCheck(servers *[]Server, ticker *time.Ticker) {
	for range ticker.C {
		for index, server := range *servers {
			healthCheckUrl := fmt.Sprintf("%s/healthcheck", server.endpointUrl)
			backendClient := &http.Client{}
			response, err := backendClient.Get(healthCheckUrl)
			if err != nil {
				fmt.Println("Issue while calling Backend")
				(*servers)[index].isHealthy = false
				continue
			}
			if response.StatusCode != 200 {
				fmt.Printf("%s is unhealthy", server.endpointUrl)
				(*servers)[index].isHealthy = false
			}
			(*servers)[index].isHealthy = true
		}
	}
}

func callBackend(backendUrl string, responseChannel chan ServerResponse) {
	backendClient := &http.Client{}

	response, err := backendClient.Get(backendUrl)
	if err != nil {
		fmt.Println("Issue while calling Backend")
		responseChannel <- ServerResponse{503, "Server unavailable"}
		return
	}
	defer response.Body.Close()
	
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in parsing backend response")
	}
	responseChannel <- ServerResponse{200, string(responseBody)}
}