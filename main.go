package main

import (
	"log"
	"net/http"

	"webapp/internal/handler"
	"webapp/internal/storage"
)

func main() {
	db, err := storage.NewPostgresDB()
	if err != nil {
		log.Fatalf("Не вдалося підключитися до бази: %s", err.Error())
	}
	defer db.Close()

	h := handler.NewHandler(db)

	// ВАЖЛИВО: должен быть h.InitRoutes()
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", h.InitRoutes()))
}
