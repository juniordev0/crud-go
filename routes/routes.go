package routes

import (
    "produtos-api/handlers"
    "net/http"

    "github.com/gorilla/mux"
)

// Configura e retorna as rotas da aplicação
func SetRoutes() *mux.Router {
    r := mux.NewRouter()

    // Rotas básicas

    r.HandleFunc("/produtos", handlers.GetProdutos).Methods("GET")
    r.HandleFunc("/produtos", handlers.CreateProduto).Methods("POST")
    r.HandleFunc("/produtos/{id}", handlers.UpdateProduto).Methods("PUT")
    r.HandleFunc("/produtos/{id}", handlers.DeleteProduto).Methods("DELETE")

    // Handlers padrão para 404 e 405
    r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
    r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowed)

    return r
}
