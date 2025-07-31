package main

import (
	"log"
	"net/http"
	"produtos-api/routes"
)

func main() {
	// Chama a função SetRoutes que retorna um roteador com todas as rotas configuradas
	r := routes.SetRoutes()

	// Inicia o servidor usando o roteador configurado
	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
