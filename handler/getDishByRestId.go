package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
)

func GetDishesByRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID := r.URL.Query().Get("restaurant_id")
	if restaurantID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "restaurant_id is required",
		})
		return
	}

	dishes, err := dbHelper.GetDishesByRestaurant(restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "failed to fetch dishes",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dishes fetched successfully",
		"data":    dishes,
	})
}
