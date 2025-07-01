package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetDishesByCreator(userID string) ([]models.AllDishReq, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	dishes := make([]models.AllDishReq, 0)
	err = database.DB.Select(&dishes, `
		SELECT id, name, price, restaurant_id, created_by, created_at
		FROM dishes
		WHERE created_by = $1 AND archived_at IS NULL
		ORDER BY created_at DESC
	`, uid)
	return dishes, err
}
