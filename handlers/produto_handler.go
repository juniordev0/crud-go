package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"produtos-api/db"
	"produtos-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

// Lista todos os produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, nome, descricao, preco, quantidade, criado_em, atualizado_em FROM produtos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []models.Produto

	for rows.Next() {
		var p models.Produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade, &p.CriadoEm, &p.AtualizadoEm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	// Retorna o JSON formatado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

// Insere um novo produto
func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var p models.Produto
	json.NewDecoder(r.Body).Decode(&p)

	err := db.DB.QueryRow(
		"INSERT INTO produtos (nome, descricao, preco, quantidade, criado_em) VALUES ($1, $2, $3, $4, NOW()) RETURNING id",
		p.Nome, p.Descricao, p.Preco, p.Quantidade,
	).Scan(&p.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Atualiza um produto
func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var p models.Produto
	json.NewDecoder(r.Body).Decode(&p)

	_, err := db.DB.Exec("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4, atualizado_em=NOW() WHERE id=$5",
		p.Nome, p.Descricao, p.Preco, p.Quantidade, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto %d atualizado com sucesso.", id)
}

// Exclui um produto
func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := db.DB.Exec("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto %d excluído com sucesso.", id)
}
