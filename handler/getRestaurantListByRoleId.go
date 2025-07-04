package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func GetMyRestaurantsByAdminId(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	restaurants, err := dbHelper.GetRestaurantsByCreator(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to get restaurant")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "restaurants fetched successfully",
		"data":    restaurants,
	})

}
