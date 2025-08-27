package handlers

import (
    "net/http"
    "produtos-api/utils"
)

// NotFound retorna 404 em JSON padronizado
func NotFound(w http.ResponseWriter, r *http.Request) {
    utils.JSONError(w, http.StatusNotFound, "rota não encontrada")
}

// MethodNotAllowed retorna 405 em JSON padronizado
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
    utils.JSONError(w, http.StatusMethodNotAllowed, "método não permitido")
}

