package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
    router := mux.NewRouter()
    router.HandleFunc("/health", HealthCheckHandler)

    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    res := httptest.NewRecorder()

    router.ServeHTTP(res, req)

    assert.Equal(t, http.StatusOK, res.Code)
    assert.Equal(t, "OK", res.Body.String())
}
