package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		HealthCheckHandler(w, r)
	})

	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
