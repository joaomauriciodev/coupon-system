# Coupon System API

API simples para gerenciamento de **cupons de desconto**, desenvolvida
em **Go** usando Gin, GORM e PostgreSQL.

------------------------------------------------------------------------

# Rodando o Projeto

## 1. Subir o banco de dados

``` bash
docker compose up -d
```

Verificar containers:

``` bash
docker ps
```

------------------------------------------------------------------------

## 2. Rodar a API

``` bash
go run cmd/server/main.go
```

Servidor ficará disponível em:

    http://localhost:8080

------------------------------------------------------------------------

# Endpoints

## Criar Cupom

Cria um novo cupom de desconto.

    POST /coupons

### Curl

``` bash
curl -X POST http://localhost:8080/coupons -H "Content-Type: application/json" -d '{
  "code": "WELCOME10",
  "discount_type": "percentage",
  "discount_value": 10,
  "max_uses": 100,
  "min_order_value": 50,
  "expires_at": "2026-12-31T00:00:00Z"
}'
```

------------------------------------------------------------------------

## Listar Cupons

Retorna todos os cupons cadastrados.

    GET /coupons

### Curl

``` bash
curl http://localhost:8080/coupons
```

------------------------------------------------------------------------

# Exemplo de Resposta

``` json
[
  {
    "id": 1,
    "code": "WELCOME10",
    "discount_type": "percentage",
    "discount_value": 10,
    "max_uses": 100,
    "current_uses": 0,
    "min_order_value": 50,
    "expires_at": "2026-12-31T00:00:00Z",
    "active": true,
    "created_at": "2026-03-05T20:00:00Z"
  }
]
```

# Tecnologias

-   Go
-   Gin
-   GORM
-   PostgreSQL
-   Docker
