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

	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(`
		INSERT INTO users (name, email, password, created_by)
		VALUES ($1, $2, $3, $4)
	`, req.Name, req.Email, string(hashedPassword), creatorID)
	if err != nil {
		return err
	}

	var userId uuid.UUID
	err = tx.Get(&userId, `SELECT id FROM users WHERE email = $1`, req.Email)
	if err != nil {
		return fmt.Errorf("failed to fetch user ID, error :: %w", err)
	}
	createdByUUID, err := uuid.Parse(creatorID)
	if err != nil {
		return err
	}

	for _, role := range req.Roles {
		var roleUUID uuid.UUID
		err = tx.Get(&roleUUID,
			`SELECT id 
				   FROM roles 
				   WHERE role_name = $1`,
			role)
		if err != nil {
			return fmt.Errorf("failed to fetch role ID, error :: %w", err)
		}
		_, err = tx.Exec(
			`INSERT INTO user_roles (user_id, role_id, created_by) 
				   VALUES ($1, $2, $3)`, userId, roleUUID, createdByUUID)
		if err != nil {
			return err
		}
	}
	return nil
}
