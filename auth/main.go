package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"users_api/api"
	"users_api/db"
	"users_api/middleware"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnectionDB() (db *sql.DB, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		return
	}

	return
}

func main() {
	conn, err := CreateConnectionDB()
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to establish connection to database: %s", err.Error())
	}

	userStore := db.CreateUserStore(conn)
	userService := middleware.CreateUserService(userStore)
	userApi := api.CreateUserApi(userService)

	router := api.InitRouter(userApi)

	srv := &http.Server{
		Handler:      router.Router(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on port 8000...")

	log.Fatal(srv.ListenAndServe())
}
