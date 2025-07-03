package dbHelper

import (
	"fmt"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetAllUsers(limit int, offset int) ([]models.SubAdminCreatedBy, error) {
	subAdmins := make([]models.SubAdminCreatedBy, 0)
	err := database.DB.Select(&subAdmins,
		`SELECT u.id, u.name, u.email, u.created_at
	    FROM users u 
		JOIN user_roles ur ON u.id = ur.user_id
		JOIN roles r ON ur.role_id = r.id
		WHERE r.role_name = 'user' 
		AND u.archived_at IS NULL
		ORDER BY u.created_at DESC
		LIMIT $1 OFFSET $2`, limit, offset)
	fmt.Println("query with pagination working...")
	return subAdmins, err
}
