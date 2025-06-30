package dbHelper

import (
	"rmssystem_1/database"
)

func DeleteSessionRecord(userID string) error {
	_, err := database.DB.DB.Exec(`
		DELETE FROM sessions
		WHERE user_id = $1
	`, userID)
	return err
}
