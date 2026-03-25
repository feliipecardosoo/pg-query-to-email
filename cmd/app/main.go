package main

import (
	"context"
	"log"

	"pg-query-to-email/internal/database"
	"pg-query-to-email/internal/repository"
)

func main() {
	// ──────────────────────────────────────────────────
	// 1. Database connection
	// ──────────────────────────────────────────────────
	dbPoolWrapper := database.NewDatabase()
	if err := dbPoolWrapper.Connect(); err != nil {
		log.Fatal(err)
	}
	defer dbPoolWrapper.Close()

	pool := dbPoolWrapper.GetPool()

	// ──────────────────────────────────────────────────
	// 2. Repository initialization
	// ──────────────────────────────────────────────────
	repo := repository.NewQueryRepository(pool)

	// ──────────────────────────────────────────────────
	// 3. Context
	// ──────────────────────────────────────────────────
	ctx := context.Background()

	// ──────────────────────────────────────────────────
	// 4. Executar query
	// ──────────────────────────────────────────────────
	users, err := repo.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ──────────────────────────────────────────────────
	// 5. Ver resultado
	// ──────────────────────────────────────────────────
	for _, u := range users {
		log.Printf("Nome: %s | Email: %s | Telefone: %s | CreatedAt: %s", u.Nome, u.Email, u.Telefone, u.CreatedAt)
	}
}
