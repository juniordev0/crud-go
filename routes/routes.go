package routes

import (
	"produtos-api/handlers"

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

	return r
}
