package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("starting the server")

	client := &http.Client{}

	rand.Seed(time.Now().UTC().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if port == "" {
		port = "localhost"
	}

	allRoutes := []string{
		"/maybe/error",
		"/maybe/fail",
		"/random/error",
	}

	route := allRoutes[0]

	for {
		randRoute := rand.Intn(len(allRoutes))
		route = allRoutes[randRoute]
		fmt.Printf(fmt.Sprintf("http://%s:%s%s : ", host, port, route))
		resp, _ := client.Get(fmt.Sprintf("http://%s:%s%s", host, port, route))
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		//time.Sleep(time.Millisecond(100000))
		time.Sleep(time.Second / 50)
	}
}
