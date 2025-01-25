package routes

import (
    "net/http"

    "movie-manage-go/api/handlers"
    "github.com/gorilla/mux"
)

// RegisterRoutes configura as rotas da aplicação
func RegisterRoutes(router *mux.Router) {
    // Rota do health check
    router.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet)
}
