package configenv

import (
	"log"

	"github.com/joho/godotenv"
)

type envStruct struct{}

func NewEnv() *envStruct {
	return &envStruct{}
}

func (e *envStruct) LoadEnv() {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
