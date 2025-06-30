package dbHelper

import (
	"rmssystem_1/database"
)

func DeleteSessionRecord(userID string) error {
	_, err := database.DB.Exec(`
		UPDATE sessions SET archived_at = NOW()
		WHERE user_id = $1;
	 `, userID)
	return err
}
