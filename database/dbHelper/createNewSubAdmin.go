package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateNewSubAdmin(req models.SubAdminReq, creator string) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	createdById, err := uuid.Parse(creator)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, req.Name, req.Email, hashedPassword)
	if err != nil {
		return err
	}
	var userID uuid.UUID
	err = database.DB.Get(&userID,
		`SELECT id 
			   FROM users 
			   WHERE email = $1`,
		req.Email)
	if err != nil {
		return err
	}

	var roleID uuid.UUID
	err = database.DB.Get(&roleID, `
		SELECT id FROM roles WHERE role_name = 'subAdmin'
	`)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO user_roles (user_id, role_id, created_by)
		VALUES ($1, $2, $3)
	`, userID, roleID, createdById)

	return err
}
