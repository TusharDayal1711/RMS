package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func GetUsersCreatedById(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	users, err := dbHelper.GetUsersCreatedBy(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to get users")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "users fetched successfully",
		"data":    users,
	})
}
