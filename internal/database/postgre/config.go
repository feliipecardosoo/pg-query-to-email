package postgre

import (
	"context"
	"fmt"
	"os"
	configenv "pg-query-to-email/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

type config struct{}

func NewConfig() *config {
	return &config{}
}

// Connect abre uma conexão com o PostgreSQL usando variáveis de ambiente.
func (c *config) Connect() error {

	carregaEnv := configenv.NewEnv()
	carregaEnv.LoadEnv()

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return fmt.Errorf("erro ao parsear config do banco: %w", err)
	}

	// Timeout para conectar
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("erro ao criar pool: %w", err)
	}

	// Testa conexão
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("erro ao conectar no PostgreSQL: %w", err)
	}

	DB = pool
	fmt.Println("🔥 PostgreSQL conectado com sucesso!")

	return nil
}

// Close fecha o pool quando a aplicação finaliza
func (c *config) Close() {
	if DB != nil {
		DB.Close()
	}
}
