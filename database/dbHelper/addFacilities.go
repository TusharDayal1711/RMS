package dbHelper

import (
	"fmt"
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
	"strings"
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
	req.Name = strings.TrimSpace(req.Name)

	restUUID, err := uuid.Parse(req.RestaurantID)
	if err != nil {
		return fmt.Errorf("invalid restaurant ID: %w", err)
	}
	createdById, err := uuid.Parse(addedBy)
	if err != nil {
		return fmt.Errorf("invalid creator ID: %w", err)
	}

	_, err = database.DB.Exec(`
		INSERT INTO dishes (name, price, restaurant_id, created_by)
		VALUES ($1, $2, $3, $4)
	`, req.Name, req.Price, restUUID, createdById)
	if err != nil {
		return fmt.Errorf("failed to insert dish: %w", err)
	}
	return nil
}
