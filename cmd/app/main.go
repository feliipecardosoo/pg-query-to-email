package main

import (
	"context"
	"log"
	"pg-query-to-email/internal/client"
	configenv "pg-query-to-email/internal/config"
	"pg-query-to-email/internal/service"
)

func main() {
	carregaEnv := configenv.NewEnv()
	carregaEnv.LoadEnv()

	//Context
	ctx := context.Background()

	// Client (HTTP)
	notifierClient := client.NewNotifierClient()

	// Service
	notifierService := service.NewNotifierService(notifierClient)

	//  Executar fluxo
	if err := notifierService.Execute(ctx); err != nil {
		log.Fatal("erro ao executar notifier:", err)
	}

	log.Println("processo finalizado com sucesso")
}
