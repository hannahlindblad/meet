# Meet accounts service

REST API for managing accounts.

## Technologies

- Golang migrate
- PostgreSQL
- Gorilla MUX

## Start developing

Install Golang(https://golang.org/doc/install)

Install [Golang migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 

Create a `.env` file with the following content

```
POSTGRES_URL="user=postgres password=postgres dbname=users port=5433 sslmode=disable"
```

Run `docker-compose up -d`

Run migrations with 

```
migrate -database "postgres://postgres:postgres@localhost:5433/accounts?sslmode=disable" -path db/migrations up
```

Run `go build`

Run `go run main.go`

The API is available on `localhost:8000`

