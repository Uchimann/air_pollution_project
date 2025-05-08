package main

import (
	"fmt"
	"log"
	"net/http"


	"github.com/uchimann/air_pollution_project/notifier/internal/sse"
)

func main(){

	http.HandleFunc("/sse", sse.sseHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}