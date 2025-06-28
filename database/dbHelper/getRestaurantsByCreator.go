package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetRestaurantsByCreator(userID string) ([]models.Restaurant, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	var restaurants []models.Restaurant
	err = database.DB.Select(&restaurants, `
		SELECT id, name, address, longitude, latitude, created_by, created_at
		FROM restaurants
		WHERE created_by = $1 AND archived_at IS NULL
		ORDER BY created_at DESC
	`, uid)

	return restaurants, err
}
