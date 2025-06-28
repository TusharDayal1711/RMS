package dbHelper

import (
	"github.com/google/uuid"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func SetUserAddress(req models.AddressReq, userID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`
		INSERT INTO address (user_id, address, longitude, latitude)
		VALUES ($1, $2, $3, $4)
	`, userUUID, req.Address, req.Longitude, req.Latitude)

	return err
}
