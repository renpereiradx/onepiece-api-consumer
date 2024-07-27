package repository

import (
	"context"

	"github.com/renpereiradx/onepiece-api-consumer/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	InsertCharacter(ctx context.Context, character *models.Character) error
	GetCharacterById(ctx context.Context, id string) (*models.Character, error)
	// GetCharacterByName(ctx context.Context, name string) (*models.Character, error)
	// UpdateCharacter(ctx context.Context, character *models.Character) error
	// DeleteCharacter(ctx context.Context, id string) error
	InsertCharacterFull(ctx context.Context, characterDetail *models.CharacterFull) error
	// ListCharacterByDevilFruit(ctx context.Context, devilFruit string) ([]*models.Character, error)
	// ListCharacter(ctx context.Context) ([]*models.Character, error)

	InsertCharacterImage(ctx context.Context, characterImage *models.CharacterImage) error
	InsertDevilFruit(ctx context.Context, devilFruit *models.DevilFruit) error
	Close() error
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func Close() error {
	return implementation.Close()
}
