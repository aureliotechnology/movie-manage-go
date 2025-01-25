package handlers

import "net/http"

// HealthCheckHandler responde com "OK"
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
