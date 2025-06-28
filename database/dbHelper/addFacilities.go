package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func CreateNewRestaurant(req models.RestaurantReq, creator string) error {
	createdBy, err := uuid.Parse(creator)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO restaurants (name, address, longitude, latitude, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`, req.Name, req.Address, req.Longitude, req.Latitude, createdBy)

	return err
} //

func AddNewDish(req models.DishReq, addedBy string) error {
	restaurantId, err := uuid.Parse(req.RestaurantID)
	if err != nil {
		return err
	}
	createdById, err := uuid.Parse(addedBy)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(`
		INSERT INTO dishes (name, price, restaurant_id, created_by)
		VALUES ($1, $2, $3, $4)
	`, req.Name, req.Price, restaurantId, createdById)
	return err
}
