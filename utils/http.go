package utils

import (
    "encoding/json"
    "net/http"
)

// ErrorResponse padroniza respostas de erro em JSON
type ErrorResponse struct {
    Error string `json:"error"`
}

// MessageResponse padroniza respostas simples com mensagem
type MessageResponse struct {
    Message string `json:"message"`
}

// JSON escreve um payload JSON com status e Content-Type adequado
func JSON(w http.ResponseWriter, status int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if v == nil {
        return
    }
    if err := json.NewEncoder(w).Encode(v); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// JSONError escreve {"error": mensagem}
func JSONError(w http.ResponseWriter, status int, msg string) {
    JSON(w, status, ErrorResponse{Error: msg})
}

// JSONMessage escreve {"message": mensagem}
func JSONMessage(w http.ResponseWriter, status int, msg string) {
    JSON(w, status, MessageResponse{Message: msg})
}

