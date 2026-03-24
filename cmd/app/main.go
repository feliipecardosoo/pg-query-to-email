package main

import (
	"log"
	database "pg-query-to-email"
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
	// pool := dbPoolWrapper.GetPool()
}
