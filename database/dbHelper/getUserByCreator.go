package dbHelper

import (
	"fmt"
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetUsersCreatedBy(createdBy string) ([]models.UserCreatedBy, error) {
	creatorID, err := uuid.Parse(createdBy)
	fmt.Println(creatorID)
	if err != nil {
		return nil, err
	}
	usersReq := make([]models.UserCreatedBy, 0)
	err = database.DB.Select(&usersReq, `
		SELECT u.id, u.name, u.email, r.role_name AS role, u.created_at
		FROM users u
		JOIN user_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		WHERE ur.created_by = $1 AND u.archived_at IS NULL
	`, creatorID)
	fmt.Println(usersReq)
	if err != nil {
		return nil, err
	}
	return usersReq, nil
}
