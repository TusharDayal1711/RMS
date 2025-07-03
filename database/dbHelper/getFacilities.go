package dbHelper

import (
	db "rmssystem_1/database"
	"rmssystem_1/models"
)

func GetAllRestaurant(limit, offset int) ([]models.RestaurantReq, error) {
	restaurants := make([]models.RestaurantReq, 0)
	err := db.DB.Select(&restaurants, `
		SELECT name, address, longitude, latitude
		FROM restaurants 
		WHERE archived_at IS NULL 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	return restaurants, err
}

func GetAllDishes(limit, offset int) ([]models.AllDishReq, error) {
	dishes := make([]models.AllDishReq, 0)
	query := `
		SELECT id, name, price, restaurant_id, created_by, created_at
		FROM dishes
		WHERE archived_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	err := db.DB.Select(&dishes, query, limit, offset)
	return dishes, err
}
