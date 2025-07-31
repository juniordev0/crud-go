package models

import (
	"produtos-api/utils"
)

// Estrutura que representa a tabela `produtos` no banco
type Produto struct {
	ID           int         `json:"id"`
	Nome         string      `json:"nome"`
	Preco        float64     `json:"preco"`
	Descricao    string      `json:"descricao"`
	Quantidade   int         `json:"quantidade"`
	CriadoEm     *utils.Date `json:"criado_em"`
	AtualizadoEm *utils.Date `json:"atualizado_em"`
}
