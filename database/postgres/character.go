package postgres

import (
	"context"
	"log"

	"github.com/renpereiradx/onepiece-api-consumer/models"
)

func (r *PostgresRepository) InsertCharacter(ctx context.Context, character *models.Character) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO characters (id, name, jname, rname) VALUES ($1, $2, $3, $4)", character.ID, character.Name, character.Jname, character.Rname)
	return err
}

func (r *PostgresRepository) InsertCharacterFull(ctx context.Context, characterDetail *models.CharacterFull) error {
	err := r.InsertCharacter(ctx, &characterDetail.Character)
	if err != nil {
		if characterDetail.CharacterImage != nil {
			err = r.InsertCharacterImage(ctx, characterDetail.CharacterImage)
			if err != nil {
				return err
			}
		}
		if characterDetail.DevilFruit != nil {
			err = r.InsertDevilFruit(ctx, characterDetail.DevilFruit)
			if err != nil {
				return err
			}
		}
		_, err = r.db.ExecContext(ctx, "INSERT INTO characters_full (id, character_id, age, birth, affiliation, ocupation, origin, height, devil_fruit_id, character_image_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", characterDetail.ID, characterDetail.Character.ID, characterDetail.Age, characterDetail.Birth, characterDetail.Affiliation, characterDetail.Ocupation, characterDetail.Origin, characterDetail.Height, characterDetail.DevilFruit.ID, characterDetail.CharacterImage.ID)
		return err
	}
	log.Println("Error inserting character: ", err)
	return err
}

func (r *PostgresRepository) InsertCharacterImage(ctx context.Context, characterImage *models.CharacterImage) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO characters_images (id, character_id, image_url) VALUES ($1, $2, $3)", characterImage.ID, characterImage.CharacterID, characterImage.ImageURL)
	return err
}

func (r *PostgresRepository) InsertDevilFruit(ctx context.Context, devilFruit *models.DevilFruit) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO devil_fruits (id, name, type) VALUES ($1, $2, $3)", devilFruit.ID, devilFruit.Name, devilFruit.Type)
	return err
}

func (r *PostgresRepository) GetCharacterById(ctx context.Context, id string) (*models.Character, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, jname, rname FROM characters WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}()
	character := &models.Character{}
	for rows.Next() {
		err = rows.Scan(&character.ID, &character.Name, &character.Jname, &character.Rname)
		if err == nil {
			return character, nil
		}
	}
	return nil, err
}
