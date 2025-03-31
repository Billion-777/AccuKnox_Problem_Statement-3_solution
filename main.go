package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	fmt.Fprintf(w, "Current Date and Time: %s\nHostname: %s", currentTime, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}