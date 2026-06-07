Markdown
# Chirpy API

## 🐦 What is it?
Chirpy is a lightweight, fast, and secure backend REST API for a Twitter-like social media application, written entirely in Go. It handles everything from user registration and secure authentication (using JWTs and refresh tokens) to creating, reading, and deleting short text posts called "Chirps." It also includes webhook integrations for processing premium user subscriptions ("Chirpy Red").

## 💡 Why should you care?
Chirpy serves as an excellent reference architecture for building modern Go web servers without relying on bloated web frameworks. It demonstrates best practices for:
* **Standard Library Routing:** Utilizing Go's native `net/http` package for clean and efficient API routing.
* **Type-Safe Database Interactions:** Using `sqlc` to generate type-safe Go code directly from raw PostgreSQL queries.
* **Database Migrations:** Managing schema evolution safely using `goose`.
* **Security First:** Implementing Argon2 password hashing and a robust, dual-token (Access + Refresh) JWT authentication system.

## 🚀 How to Install and Run

### Prerequisites
* [Go](https://golang.org/doc/install) (1.22 or higher recommended)
* [PostgreSQL](https://www.postgresql.org/download/)
* [Goose](https://github.com/pressly/goose) (for database migrations)
* [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) (optional, only needed if you want to modify SQL queries)

### 1. Setup your Environment
Clone the repository and create a `.env` file in the root directory with the following variables:

```env
DB_URL=postgres://your_username:your_password@localhost:5432/chirpy?sslmode=disable
PLATFORM=dev
JWT_SECRET=generate_a_random_64_character_string
POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e
(Tip: You can generate a JWT secret in your terminal using openssl rand -base64 64)

2. Initialize the Database
Ensure your PostgreSQL server is running and create a database named chirpy. Then, apply the schema migrations:

Bash
cd sql/schema
goose postgres "postgres://your_username:your_password@localhost:5432/chirpy?sslmode=disable" up
cd ../../
3. Build and Run
Download the Go modules, build the executable, and start the server:

Bash
# Download dependencies
go mod download

# Build and run the server
go build -o out && ./out
The server will start running locally on port 8080. You can test it by visiting http://localhost:8080/api/healthz.