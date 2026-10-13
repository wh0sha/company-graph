package main

import (
	"company-graph/internal/storage"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	connString := "postgres://user:password@localhost:5432/companies?sslmode=disable"

	store, err := storage.NewStorage(ctx, connString)
	if err != nil {
		log.Fatalf("Не удалось подлючиться к бд: %v", err)
	}
	fmt.Println("бд подключена")

	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "метод не разрешон", http.StatusMethodNotAllowed)
		}

		companies, err := store.GetCompanies(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(companies)
		if err != nil {
			log.Printf("ошибка при отправке")
		}

	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("ошибка запуска сервака: %v", err)
	}
}
