package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func main() {
	router := ddd.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
