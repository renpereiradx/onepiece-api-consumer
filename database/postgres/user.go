package postgres

import (
	"context"
	"log"

	"github.com/renpereiradx/onepiece-api-consumer/models"
)

func (r *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password)
	return err
}

func (r *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	user := models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, err
}
