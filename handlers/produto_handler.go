package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "produtos-api/db"
    "produtos-api/models"
    "produtos-api/utils"
    "strconv"

    "github.com/gorilla/mux"
)

// Lista todos os produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, nome, descricao, preco, quantidade, criado_em, atualizado_em FROM produtos")
    if err != nil {
        utils.JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }
    defer rows.Close()

	var produtos []models.Produto

	for rows.Next() {
		var p models.Produto
        if err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade, &p.CriadoEm, &p.AtualizadoEm); err != nil {
            utils.JSONError(w, http.StatusInternalServerError, err.Error())
            return
        }
        produtos = append(produtos, p)
    }

    utils.JSON(w, http.StatusOK, produtos)
}

// Insere um novo produto
func CreateProduto(w http.ResponseWriter, r *http.Request) {
    var p models.Produto
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        utils.JSONError(w, http.StatusBadRequest, "JSON inválido")
        return
    }

    err := db.DB.QueryRow(
        "INSERT INTO produtos (nome, descricao, preco, quantidade, criado_em) VALUES ($1, $2, $3, $4, NOW()) RETURNING id",
        p.Nome, p.Descricao, p.Preco, p.Quantidade,
    ).Scan(&p.ID)

    if err != nil {
        utils.JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.JSON(w, http.StatusCreated, p)
}

// Atualiza um produto
func UpdateProduto(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.JSONError(w, http.StatusBadRequest, "id inválido")
        return
    }

    var p models.Produto
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        utils.JSONError(w, http.StatusBadRequest, "JSON inválido")
        return
    }

    res, err := db.DB.Exec("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4, atualizado_em=NOW() WHERE id=$5",
        p.Nome, p.Descricao, p.Preco, p.Quantidade, id)

    if err != nil {
        utils.JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }

    if rows, _ := res.RowsAffected(); rows == 0 {
        utils.JSONError(w, http.StatusNotFound, "produto não encontrado")
        return
    }

    utils.JSONMessage(w, http.StatusOK, fmt.Sprintf("Produto %d atualizado com sucesso.", id))
}

// Exclui um produto
func DeleteProduto(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.JSONError(w, http.StatusBadRequest, "id inválido")
        return
    }

    res, err := db.DB.Exec("DELETE FROM produtos WHERE id = $1", id)
    if err != nil {
        utils.JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }

    if rows, _ := res.RowsAffected(); rows == 0 {
        utils.JSONError(w, http.StatusNotFound, "produto não encontrado")
        return
    }

    // 204 No Content, sem corpo
    w.WriteHeader(http.StatusNoContent)
}
