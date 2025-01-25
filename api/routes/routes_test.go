package routes

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
    // Configurar o roteador
    router := mux.NewRouter()
    RegisterRoutes(router)

    // Testar a rota de health check
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    res := httptest.NewRecorder()

    router.ServeHTTP(res, req)

    // Verificar a resposta
    assert.Equal(t, http.StatusOK, res.Code)
    assert.Equal(t, "OK", res.Body.String())
}
