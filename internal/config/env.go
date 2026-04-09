package configenv

import (
	"log"

	"github.com/joho/godotenv"
)

// envStruct é a estrutura que representa a configuração do ambiente. Ela pode ser expandida no futuro para incluir mais campos ou métodos relacionados à configuração.
type envStruct struct{}

// NewEnv é um construtor para envStruct, que retorna uma nova instância da estrutura. Ele pode ser utilizado para criar um objeto de configuração do ambiente e chamar o método LoadEnv para carregar as variáveis de ambiente a partir do arquivo .env.
func NewEnv() *envStruct {
	return &envStruct{}
}

// LoadEnv é o método responsável por carregar as variáveis de ambiente a partir do arquivo .env localizado no caminho especificado. Ele utiliza a biblioteca godotenv para ler o arquivo e carregar as variáveis de ambiente no processo. Caso ocorra algum erro ao carregar o arquivo, ele loga uma mensagem de erro e encerra a aplicação.
func (e *envStruct) LoadEnv() {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
