package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

// { x: 10, y: 20 }
type RequestStruct struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ResponseStruct struct {
	Result float64 `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var request RequestStruct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := math.Max(request.X, request.Y)

	json.NewEncoder(w).Encode(ResponseStruct{Result: result + 1})
}

func main() {
	port := 8080
	host := "0.0.0.0"
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Println("Starting server on localhost:8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
