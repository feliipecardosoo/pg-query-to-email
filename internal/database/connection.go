package database

import (
	"pg-query-to-email/internal/database/postgre"

	"github.com/jackc/pgx/v5/pgxpool"
)

type databaseStruct struct{}

func NewDatabase() *databaseStruct {
	return &databaseStruct{}
}

func (d *databaseStruct) Connect() error {
	conexaoDb := postgre.NewConfig()
	return conexaoDb.Connect()
}

func (d *databaseStruct) GetPool() *pgxpool.Pool {
	return postgre.DB
}

func (d *databaseStruct) Close() {
	closeConexaoDb := postgre.NewConfig()
	closeConexaoDb.Close()
}
