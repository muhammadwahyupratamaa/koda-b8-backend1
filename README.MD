# Backend Authentication API

Backend sederhana menggunakan **Golang**, **Gin**, dan **PostgreSQL** untuk implementasi fitur **Register**, **Login**, dan **Get All Users**.

## Features

- Register User
- Login User
- Get All Users
- PostgreSQL Database
- Repository Pattern
- Service Layer
- Dependency Injection

---

## Project Structure

```
.
├── cmd
├── internal
│   ├── di
│   ├── handler
│   ├── lib
│   ├── model
│   ├── repo
│   └── svc
├── migrations
├── init.sql
├── request.http
└── README.md
```

---

## Application Flow

```
Client
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
PostgreSQL
```

---

## Installation

Clone repository

```bash
git clone https://github.com/username/koda-b8-backend1.git
```

Masuk ke folder project

```bash
cd koda-b8-backend1
```

Install dependency

```bash
go mod tidy
```

---

## Environment Variables

Buat file `.env`

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/authentication?sslmode=disable
```

---

## Menjalankan Project

```bash
go run ./cmd/main.go
```

atau jika menggunakan Air

```bash
air
```

Server akan berjalan di

```
http://localhost:8080
```

---

## API Endpoint

| Method | Endpoint | Keterangan |
|---------|----------|------------|
| POST | /register | Register User |
| POST | /login | Login User |
| GET | /users | Menampilkan seluruh data user |

---

## Testing Endpoint

Project ini menggunakan **REST Client** pada VS Code.

### Register

```http
POST http://localhost:8080/register
Content-Type: application/x-www-form-urlencoded

name=wahyu&email=wahyu@mail.com&password=12345678
```

### Login

```http
POST http://localhost:8080/login
Content-Type: application/x-www-form-urlencoded

email=wahyu@mail.com&password=12345678
```

### Get All Users

```http
GET http://localhost:8080/users
```

---

## Database Schema

```mermaid
erDiagram

users {
BIGSERIAL id PK
VARCHAR name
VARCHAR email UK
VARCHAR password
TIMESTAMP created_at
TIMESTAMP updated_at
}
```

---

## Author

Muhammad Wahyu Pratama