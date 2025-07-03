package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func GetRestaurantByID(restaurantID string) (models.RestaurantReq, error) {
	restUUID, err := uuid.Parse(restaurantID)
	if err != nil {
		return models.RestaurantReq{}, err
	}
	var restRes models.RestaurantReq
	err = database.DB.Get(&restRes, `
		SELECT name, address, longitude, latitude 
		FROM restaurants 
		WHERE id = $1`, restUUID)
	return restRes, err
}
