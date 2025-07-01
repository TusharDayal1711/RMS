package dbHelper

import (
	db "rmssystem_1/database"
	"rmssystem_1/models"
)

func GetaAllRestaurant() ([]models.RestaurantReq, error) {
	restaurants := make([]models.RestaurantReq, 0)
	query := `SELECT id, name, address FROM restaurants WHERE archived_at IS NULL`
	err := db.DB.Select(&restaurants, query)
	return restaurants, err
}

func GetAllDishes() ([]models.AllDishReq, error) {
	dishes := make([]models.AllDishReq, 0)
	err := db.DB.Select(&dishes, `
		SELECT id, name, price, restaurant_id, created_by, created_at
		FROM dishes
		WHERE archived_at IS NULL
		ORDER BY created_at DESC
	`)
	return dishes, err
}
