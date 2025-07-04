package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
)

func GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	restaurantID := r.URL.Query().Get("restaurant_id")
	if restaurantID == "" {
		http.Error(w, "restaurant_id is required", http.StatusBadRequest)
		return
	}
	restaurant, err := dbHelper.GetRestaurantByID(restaurantID)
	if err != nil {
		http.Error(w, "restaurant not found or invalid ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurant)
}
