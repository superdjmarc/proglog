package main

import (
	"log"

	"github.com/superdjmarc/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
