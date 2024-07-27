package repository

import (
	"context"

	"github.com/renpereiradx/onepiece-api-consumer/models"
)

func InsertCharacter(ctx context.Context, character *models.Character) error {
	return implementation.InsertCharacter(ctx, character)
}

func GetCharacterById(ctx context.Context, id string) (*models.Character, error) {
	return implementation.GetCharacterById(ctx, id)
}

// func GetCharacterByName(ctx context.Context, name string) (*models.Character, error) {
// 	return implementation.GetCharacterByName(ctx, name)
// }

// func UpdateCharacter(ctx context.Context, character *models.Character) error {
// 	return implementation.UpdateCharacter(ctx, character)
// }

// func DeleteCharacter(ctx context.Context, id string) error {
// 	return implementation.DeleteCharacter(ctx, id)
// }

// func ListCharacterByDevilFruit(ctx context.Context, devilFruit string) ([]*models.Character, error) {
// 	return implementation.ListCharacterByDevilFruit(ctx, devilFruit)
// }

// func ListCharacter(ctx context.Context) ([]*models.Character, error) {
// 	return implementation.ListCharacter(ctx)
// }
