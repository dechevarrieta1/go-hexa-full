package postgresadapter

import (
	"database/sql"
	"fmt"
	usermodels "go-hexa-full/internal/core/user/models"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(connString string) (*PostgresUserRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %w", err)
	}

	return &PostgresUserRepository{db: db}, nil
}

func (r *PostgresUserRepository) Save(user usermodels.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET name=$2", user.Name, user.Name)
	return err
}

func (r *PostgresUserRepository) FindById(id string) (*usermodels.User, error) {
	var user usermodels.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id=$1", id).Scan(&user.Name, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
