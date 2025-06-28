package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetUsersCreatedBy(creatorBy string) ([]models.UserCreatedBy, error) {
	CreatorId, err := uuid.Parse(creatorBy)
	if err != nil {
		return nil, err
	}

	var users []models.UserCreatedBy
	err = database.DB.Select(&users, `
		SELECT u.id, u.name, u.email, r.role_name as role, u.created_at
		FROM users u
		JOIN user_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		WHERE ur.created_by = $1 AND u.archived_at IS NULL
	`, CreatorId)
	return users, err
}
