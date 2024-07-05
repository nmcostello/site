package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nmcostello/site/handlers"
	"github.com/nmcostello/site/session"
	"golang.org/x/exp/slog"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	h := handlers.New(log)

	var secureFlag = true
	if os.Getenv("SECURE_FLAG") == "false" {
		secureFlag = false
	}

	// Add session middleware.
	sh := session.NewMiddleware(h, session.WithSecure(secureFlag))

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      sh,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on %v\n", server.Addr)
	log.Error("Error starting server", server.ListenAndServe())
}
