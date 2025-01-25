package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
    // Arrange: Configuração do handler
    handler := http.HandlerFunc(HealthCheckHandler)

    // Act: Simula uma requisição HTTP GET para o endpoint
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    res := httptest.NewRecorder()

    handler.ServeHTTP(res, req)

    // Assert: Verifica se a resposta está como esperado
    assert.Equal(t, http.StatusOK, res.Code)
    assert.Equal(t, "OK", res.Body.String())
}
