package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Zhang-Yu-Bo/friendly-pancake/router"
	"github.com/gorilla/handlers"
)

func main() {
	port := "80"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	mRouter := router.NewRouter()

	origins := strings.Split(os.Getenv("ORIGIN_ALLOWED"), ",")
	if len(origins) == 1 && origins[0] == "" {
		origins[0] = "*"
	}
	originsOk := handlers.AllowedOrigins(origins)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	mServer := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(mRouter),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server run on port:", port)

	log.Fatal(mServer.ListenAndServe())
}
