package main

import (
	"log"
	"net/http"

	"Credis/platform/authenticator"
	"Credis/platform/router"
	"Credis/web/app/db"

	"github.com/joho/godotenv"
)

func main() {
	db.Connect()
	defer db.Disconnect()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("HTTP server is listening on port 192:3000...")

	if err := http.ListenAndServe("192.168.56.1:3000", rtr); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

}
