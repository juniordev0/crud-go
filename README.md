# 📦 Projeto de API em Go - Produtos

Este projeto é uma API RESTful desenvolvida em **Go (Golang)** para gerenciar produtos em um banco de dados PostgreSQL. Ele permite criar, listar, atualizar e deletar produtos.

## 🚀 Tecnologias utilizadas

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) (opcional)
- `net/http` para criação das rotas
- `database/sql` para conexão com o banco
- Organização em packages: `models`, `utils`, `controllers`, `db`

## 📁 Estrutura de Pastas

├── controllers/ # Lógica dos endpoints
├── db/          # Conexão com banco de dados
├── models/      # Definição de structs e interfaces
├── routes/      # Rotas da API
├── utils/       # Funções auxiliares e formatadores
├── postgreSQL/  # Contém a DDL da tabela produtos
├── main.go      # Ponto de entrada da aplicação
└── go.mod / go.sum # Gerenciamento de dependências

## 🔧 Instalação e Execução

### Pré-requisitos

- Go instalado (versão 1.18 ou superior)
- PostgreSQL rodando
- (Opcional) Docker e Docker Compose

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/nome-do-repo.git
cd nome-do-repo
```

### 2. Crie o banco de dados
- `CREATE DATABASE postgres;`
- Na pasta `postgreSQL` existe a DDL para criação da tabela.

### 3. Configure a conexão com o banco
- No arquivo `db/connection.go`, configure os dados de conexão com o banco:
  - `usuario := "postgres"`
  - `senha := "sua_senha"`
  - `host := "localhost"`
  - `porta := "5432"`
  - `dbname := "produtos_db"`

### 4. Instale as dependências
- `go mod tidy`

### 5. Execute o servidor
- `go run main.go`

A API ficará disponível em:
📍 http://localhost:8080

### 🛠️ Endpoints disponíveis

| Método | Endpoint       | Descrição               |
| ------ | -------------- | ----------------------- |
| GET    | /produtos      | Lista todos os produtos |
| POST   | /produtos      | Cria um novo produto    |
| PUT    | /produtos/{id} | Atualiza um produto     |
| DELETE | /produtos/{id} | Deleta um produto       |

📌 **Observações**
- Campos de data podem ser enviados como `YYYY-MM-DD` ou `MM-DD-YYYY` e são retornados como `YYYY-MM-DD`.
- O código segue princípios de separação de responsabilidades.

📄 Licença
Este projeto está licenciado sob a MIT License.
