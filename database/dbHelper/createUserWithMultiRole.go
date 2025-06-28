package dbHelper

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func CreateUserWithMultiRole(req models.MultiRole, creatorID string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`, req.Name, req.Email, string(hashedPassword))
	if err != nil {
		return err
	}

	var userId uuid.UUID
	err = database.DB.Get(&userId, `SELECT id FROM users WHERE email = $1`, req.Email)
	if err != nil {
		return fmt.Errorf("failed to fetch user ID, error :: %w", err)
	}
	createdByUUID, err := uuid.Parse(creatorID)
	if err != nil {
		return err
	}

	for _, role := range req.Roles {
		var roleUUID uuid.UUID
		err = database.DB.Get(&roleUUID,
			`SELECT id 
				   FROM roles 
				   WHERE role_name = $1`,
			role)
		if err != nil {
			return fmt.Errorf("failed to fetch role ID, error :: %w", err)
		}
		_, err = database.DB.Exec(
			`INSERT INTO user_roles (user_id, role_id, created_by) 
				   VALUES ($1, $2, $3)`, userId, roleUUID, createdByUUID)
		if err != nil {
			return err
		}
	}
	return nil
}
