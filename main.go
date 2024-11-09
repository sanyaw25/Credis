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

	// Start an HTTP server to redirect all HTTP requests to HTTPS
	go func() {
		httpRouter := http.NewServeMux()
		httpRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Redirect HTTP to HTTPS
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		})
		if err := http.ListenAndServe(":80", httpRouter); err != nil {
			log.Fatalf("Failed to start HTTP redirect server: %v", err)
		}
	}()

	// Start the HTTPS server
	log.Print("Server listening on https://localhost:443/")
	if err := rtr.RunTLS(":443", "server.crt", "server.key"); err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}
