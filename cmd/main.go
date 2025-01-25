package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "movie-manage-go/api/routes"
)

func main() {
    router := mux.NewRouter()

    routes.RegisterRoutes(router)

    log.Println("Servidor iniciado na porta 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
