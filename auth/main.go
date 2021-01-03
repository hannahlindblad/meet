package main

import (
	"auth/api"
	"auth/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	return
}

func createConnectionDB() (db *sql.DB, err error) {
	fmt.Println(os.Getenv("POSTGRES_URL"))
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
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment: %s", err.Error())
	}

	conn, err := createConnectionDB()

	if err != nil {
		log.Fatalf("Failed to establish connection to database: %s", err.Error())
	}

	defer conn.Close()

	userStore := db.CreateUserStore(conn)
	authApi := api.CreateAuthApi(userStore)

	router := api.InitRouter(authApi)

	srv := &http.Server{
		Handler:      router.Router(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on port 8000...")

	log.Fatal(srv.ListenAndServe())
}
