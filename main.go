package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"

	"todo/domain"
	"todo/handlers"
	"todo/postgres"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "todo_dev",
	})

	defer DB.Close()

	domainDB := domain.DB{
		UserRepo: postgres.NewUserRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}
}
