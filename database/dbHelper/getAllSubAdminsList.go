package dbHelper

import (
	"fmt"
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetAllSubAdminsList(AdminId string, limit int, offset int) ([]models.SubAdminCreatedBy, error) {
	AId, err := uuid.Parse(AdminId)
	subAdmins := make([]models.SubAdminCreatedBy, 0)
	if err != nil {
		return nil, err
	}
	err = database.DB.Select(&subAdmins,
		`SELECT u.id, u.name, u.email, u.created_at
	    FROM users u 
		JOIN user_roles ur ON u.id = ur.user_id
		JOIN roles r ON ur.role_id = r.id
		WHERE r.role_name = 'subAdmin' 
		AND ur.created_by = $1 
		AND u.archived_at IS NULL
		ORDER BY u.created_at DESC
		LIMIT $2 OFFSET $3`, AId, limit, offset)
	fmt.Println("query with pagination working...")
	return subAdmins, err
}
