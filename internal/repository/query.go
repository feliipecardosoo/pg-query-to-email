package repository

const (
	GetUsersQuery = `
        SELECT 
		nome, email, telefone, created_at
		FROM users;
    `
)
