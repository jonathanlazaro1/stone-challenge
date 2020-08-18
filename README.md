# Stone Challenge

Esta API Rest foi desenvolvida para o desafio técnico da empresa [Stone](https://www.stone.com.br), e consiste em um gerenciador de Notas Fiscais. A mesma encontra-se publicada, e sua documentação de uso pode ser consultada através do seguinte link:
https://stone-invoice-api.herokuapp.com/swagger/index.html

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

   git clone https://github.com/jonathanlazaro1/stone-challenge.git
   cd ./stone-challenge

2. Aponte o terminal para a pasta do projeto

docker-compose -f ./docker-compose.yml up --build

docker-compose -f ./docker-compose.pgsql.yml up -d
