# API Backend SaaS (Golang)

Backend em **Golang** com arquitetura SaaS, pronto para consumo por **frontends web e mobile**, utilizando **JWT**, **RBAC**, **multi-tenant**, **audit log**, **PostgreSQL** e **Docker**.

---

## 🚀 Tecnologias

- Go (Gin)
- PostgreSQL
- GORM
- JWT (golang-jwt)
- RBAC (Roles & Permissions no banco)
- Docker & Docker Compose
- Godotenv

---

## 📂 Estrutura do Projeto

```
api-backend-saas/
├── cmd/server/main.go
├── internal/
│   ├── config/
│   ├── database/
│   │   ├── database.go
│   │   └── seed.go
│   ├── handlers/
│   ├── middlewares/
│   │   ├── jwt_middleware.go
│   │   ├── permission_middleware.go
│   │   ├── tenant_middleware.go
│   │   └── audit_middleware.go
│   ├── models/
│   ├── security/
│   └── services/
├── docker-compose.yml
├── go.mod
├── go.sum
├── .env.stage
└── README.md
```

---

## ⚙️ Configuração de Ambiente

### Variáveis de Ambiente

Crie o arquivo **.env.stage** na raiz do projeto:

```env
APP_ENV=stage
SERVER_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=saas

JWT_SECRET=stage-super-secret-change-me
JWT_ISSUER=api-backend-saas
JWT_EXPIRE_MINUTES=60
```

> ⚠️ Em produção use `.env.production` com segredo forte.

---

## 🐳 Subindo o Banco com Docker

```bash
docker-compose up -d
```

PostgreSQL ficará disponível em:
```
localhost:5432
```

---

## ▶️ Rodando a API

```bash
go run cmd/server/main.go
```

A API sobe em:
```
http://localhost:8080
```

---

## 🧪 Health Check

```bash
curl http://localhost:8080/health
```

Resposta:
```json
{ "status": "ok" }
```

---

## 🔐 Autenticação

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@admin.com",
    "password": "admin123"
  }'
```

Resposta:
```json
{
  "token": "JWT_TOKEN"
}
```

---

## 👤 Criar Usuário (RBAC + JWT)

Endpoint protegido por:
- JWT
- Tenant
- Permissão `user:create`

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@empresa.com",
    "password": "senha123",
    "role_id": 2
  }'
```

---

## 🧠 Arquitetura de Segurança

- **JWT** → Autenticação
- **RBAC** → Permissões no banco
- **Multi-tenant** → `company_id` no JWT
- **Audit Log** → Todas as ações protegidas

Fluxo:
```
JWT → Tenant → Audit → Permission
```

---

## 🗄️ Seed Inicial

Ao iniciar a aplicação:
- Company padrão
- Usuário admin
- Roles
- Permissions

Tudo é criado automaticamente via `AutoMigrate + Seed`.

---

## 📌 Próximos Passos

- Listagem de usuários
- Swagger com JWT Authorize
- Refresh Token
- Soft Delete
- Testes automatizados

---

## 🧑‍💻 Autor

Projeto em evolução para uso **SaaS real**, com foco em boas práticas de backend, segurança e escalabilidade.

