package model

import "time"

type User struct {
	Nome, Email, Telefone string
	CreatedAt             time.Time
}
