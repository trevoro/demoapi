package main

import (
	"fmt"
	"os"

	"github.com/trevoro/demoapi/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}
	server := http.Server{
		Addr: fmt.Sprintf("%v:%v", "127.0.0.1", port),
	}
	server.Routes()
	server.Start()
}
