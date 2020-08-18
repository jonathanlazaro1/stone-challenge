# Stone Challenge

[![Build Status](https://dev.azure.com/jonathanlazaro1/Invoice%20API/_apis/build/status/jonathanlazaro1.stone-challenge?branchName=master)](https://dev.azure.com/jonathanlazaro1/Invoice%20API/_build/latest?definitionId=2&branchName=master)

Esta API Rest foi desenvolvida para o desafio técnico da empresa [Stone](https://www.stone.com.br), e consiste em um gerenciador de Notas Fiscais. A mesma encontra-se publicada, e sua documentação de uso pode ser consultada através do seguinte link:
https://stone-invoice-api.herokuapp.com

A API utiliza-se das seguintes tecnologias:

- Go v1.13
- Gorilla Mux (router)
- PostgreSQL v12
  - pq (driver PGSQL nativo em Go)
  - goqu (builder de queries SQL)
- Swagger (utilizando o package swag)
- Heroku, GitHub Actions

## Execução

### Máquina local

1. Clone o repositório para sua máquina e aponte o terminal para a pasta do projeto

```sh
   git clone https://github.com/jonathanlazaro1/stone-challenge.git
   cd ./stone-challenge
```

2. Se desejar utilizar o container Docker do PGSQL, basta executar o comando abaixo (necessário docker-compose):

```sh
   docker-compose -f ./docker-compose.pgsql.yml up -d
```

3. Crie o arquivo de configuração **.env** dentro da raiz do projeto, e configure-o conforme seu ambiente (banco de dados, chave de encriptação do token, etc):

```
DB_USER=<<usuario_do_bd>>
DB_PASS=<<senha_do_bd>>
DB_HOST=<<host_do_bd>>
DB_PORT=<<porta_do_bd>>
DB_NAME=<<nome_do_bd>>
DB_SSL_MODE=<<disable|require|verify-ca|verify-full>>
PORT=<<porta_da_api>>
APP_AUTH_SECRET=<<chave_de_encrypt_token>>
```

**OBS**: _há um arquivo ".env.example" que já vem pronto com as configurações do banco de dados do Docker, bastando renomeá-lo para ".env"._

4. Execute o script de migração do banco de dados:

```go
go run ./cmd/migrate/main.go
```

5.  Agora, basta rodar a aplicação e acessar a documentação via http://localhost:xxxx (onde "xxxx" é a porta da API definida no arquivo ".env"):

```go
go run ./main.go
```

### Docker

1. Clone o repositório para sua máquina e aponte o terminal para a pasta do projeto

```sh
   git clone https://github.com/jonathanlazaro1/stone-challenge.git
   cd ./stone-challenge
```

2. Execute o comando a seguir:

```sh
docker-compose -f ./docker-compose.yml up --build
```

A aplicação deverá estar disponível na porta 8080.

## Testes

Os testes foram escritos usando as ferramentas nativas do Go.
Para testes unitários, execute:

```go
go test ./domain ./helpers
```

Para testes de integração, execute:

```go
go test ./infra/handler
```

Para rodar todos os testes, execute:

```go
go test ./...
```
