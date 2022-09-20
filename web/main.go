package main

import (
	"fmt"
	"net/http"

	"github.com/muhaidil13/web-golang/pkg/handlers"
)

const portNum = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	fmt.Println("Listen on Port ", portNum)
	http.ListenAndServe(portNum, nil)
}
