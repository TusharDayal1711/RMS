package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetDishesByRestaurant(restaurantID string, limit int, offset int) ([]models.DishReq, error) {
	restUUID, err := uuid.Parse(restaurantID)
	if err != nil {
		return nil, err
	}
	dishes := make([]models.DishReq, 0)
	err = database.DB.Select(&dishes, `
		SELECT id, name, price, restaurant_id, created_by, created_at
		FROM dishes
		WHERE restaurant_id = $1 AND archived_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, restUUID, limit, offset)
	return dishes, err
}
