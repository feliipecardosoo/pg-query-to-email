package repository

const (
	GetUsersQuery = `
        SELECT 
            nome, email, telefone, created_at
        FROM users
        WHERE created_at >= $1
          AND created_at < $2;
    `
)
