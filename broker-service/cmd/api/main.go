package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("starting broker service on port %s\n", webPort)

	sev := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := sev.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
