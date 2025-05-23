# RESTful API Quote service


---

## Endpoints

- GET /quotes — retrieve all quotes. 

- GET /quotes?author={author} — retrieve quotes filtered by author. 

- GET /quotes/random — retrieve a single random quote. 

- POST /quotes — create a new quote. Provide JSON payload with quote and author. 

- DELETE /quotes/{id} — soft-delete a quote by its UUID. 

---

## Prerequisites

- Go 1.20+
- Docker + Docker Compose

---

## Getting Started

1. **Clone the repo**
   
  ```bash
  git clone https://github.com/superdumb33/quote_service_test.git
  cd quote_service_test
  ```
   
2. **Copy environment file**
   ```bash
   cp .env.example .env
   ```

3. **Start Docker services**
   ```bash
   docker-compose up --build
   ```
   
   **Line Ending Notice**
   
   Please pay attention that after cloning the repo, entrypoint.sh script may have DOS (CRLF) line endings, and Docker can only interprete Unix (LF) line endings
   
   **Quick fix**
   ```bash
   # dos2unix (if installed)
   dos2unix entrypoint.sh

   # with sed (no extra dependencies)
   sed -i 's/\r$//' entrypoint.sh
   ```
   **Or just switch it inside IDE (vscode example):**


   ![image](https://github.com/user-attachments/assets/7bffb4a9-c0a5-453c-9102-162db1449547)


  
   The service listens on **localhost:8080** by default, SwaggerUI is available at **http://localhost:8080/swagger/index.html**

---

## Configuration

Example `.env.example`:

```dotenv
#postgres-db
POSTGRES_USER=postgres
POSTGRES_DB=test_db
POSTGRES_PASSWORD=super_secret
POSTGRES_HOST=postgres
POSTGRES_PORT=5432

#app
APP_PORT=8080
```

---

## API Reference

### POST /quotes

Create a new quote.

- **Request Body** (application/json):
  - `author` (string, required)
  - `quote` (string, required)

**Example**:

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated."
}'
```

**Response (201 Created)**:

```json
{
  "id": "<uuid>"
}
```

---

### GET /quotes

Retrieve all quotes, or filter by author.

- **Query Param** (optional):
  - `author` (string) — filter results by author

**Examples**:

```bash
curl http://localhost:8080/quotes
curl http://localhost:8080/quotes?author=Confucius
```

**Response (200 OK)**:

```json
[
  {
    "id": "<uuid>",
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated.",
    "created_at": "2025-05-23T12:34:56Z"
  },
]
```

---

### GET /quotes/random

Get a single random quote.

**Example**:

```bash
curl http://localhost:8080/quotes/random
```

**Response (200 OK)**:

```json
{
  "id": "<uuid>",
  "author": "<author>",
  "quote": "<text>",
  "created_at": "<timestamp>"
}
```

**Error (404 Not Found)**:

```json
{
  "error": "Not found"
}
```

---

### DELETE /quotes/{id}

Soft-delete a quote by its UUID.

- **Path Param**:
  - `id` (UUID, required)

**Example**:

```bash
curl -X DELETE http://localhost:8080/quotes/<uuid>
```

**Response (204 No Content)**

---
**P.S.**

сервисный слой получился слишком тонким, поэтому не придумал как нормально юнит-тесты накидать =)
