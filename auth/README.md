# Meet auth service

REST API for managing users and authentication.

## Technologies

- Golang migrate
- PostgreSQL
- Gorilla MUX

## Start developing

Install [Golang migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) 

Create a `.env` file according with the following content

```
POSTGRES_URL="user=postgres password=postgres dbname=users port=5433 sslmode=disable"
```

Run `docker-compose up -d`

Run migrations with 

```
migrate -database "postgres://postgres:postgres@localhost:5433/users?sslmode=disable" -path db/migrations up
```

Run `go build`

Run `go run main.go`

The API is available on `localhost:8000`
