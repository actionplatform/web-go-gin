package main

import (
	"log"
	"os"

	"github.com/actionplatform/web-go-gin/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := ":" + port
	r := app.New()
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
