package main

import (
	"github.com/joho/godotenv"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/platform/router"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
