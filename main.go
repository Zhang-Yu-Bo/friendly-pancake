package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Zhang-Yu-Bo/friendly-pancake/router"
)

func main() {
	port := "80"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	mRouter := router.NewRouter()

	mServer := &http.Server{
		Handler:      mRouter,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server run on port:", port)

	log.Fatal(mServer.ListenAndServe())
}
