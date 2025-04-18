# Billing Prices

`billing-prices` é um serviço responsável por gerenciar preços e publicar eventos relacionados a preços no Kafka. Ele utiliza PostgreSQL como banco de dados e Kafka para comunicação assíncrona.

---

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Banco de Dados:** PostgreSQL
- **Mensageria:** Kafka
- **Ferramentas:** Docker, Docker Compose

---

## 🚀 Como Executar o Projeto

### Pré-requisitos

- [Go](https://go.dev/) (versão 1.20 ou superior)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Passos para rodar o projeto

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/luizfelipe94/billing-prices.git
   cd billing-prices