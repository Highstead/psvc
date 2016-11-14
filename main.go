package main

import (
    "log"
    "net/http"

    "app/ddd"
)

func main() {
    router := ddd.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
