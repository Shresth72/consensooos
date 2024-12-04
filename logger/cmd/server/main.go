package main

import (
	"log"

	"github.com/Shresth72/dsys/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8069")
	log.Fatal(srv.ListenAndServe())
}
