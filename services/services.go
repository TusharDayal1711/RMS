package services

import (
	"fmt"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func CalculateRestaurantDistanceByID(restaurantID, addressID string) (float64, error) {
	var RestCoord models.CoordinatesReq
	err := database.DB.Get(&RestCoord, `
		SELECT longitude, latitude
		FROM restaurants
		WHERE id = $1
	`, restaurantID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch restaurant: %w", err)
	}

	var address models.AddCoord
	err = database.DB.Get(&address, `
		SELECT longitude, latitude
		FROM address
		WHERE id = $1
	`, addressID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch address: %w", err)
	}
	ResultDistance, err := haversine(address.Latitude, address.Longitude, RestCoord.Latitude, RestCoord.Longitude), nil
	return ResultDistance, nil
}
