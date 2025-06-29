package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func GetMyRestaurantsByCreatorId(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	restaurants, err := dbHelper.GetRestaurantsByCreator(userID)
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "failed to fetch restaurants",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "restaurants fetched successfully",
		"data":    restaurants,
	})

}
