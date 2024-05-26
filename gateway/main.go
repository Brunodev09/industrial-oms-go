package main

import (
	"log"
	"net/http"

	common "github.com/brunodev09/industrial-oms-go/commons"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.Env("HTTP_ADDR", ":3000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("HTTP Server listening at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal(err)
	}
}
