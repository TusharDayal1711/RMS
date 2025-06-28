package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/utils"
)

// getting all the restaurent form db
func ListAllRestaurants(w http.ResponseWriter, r *http.Request) {
	restaurants, err := dbHelper.GetaAllRestaurant()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch restaurants")
		return
	}
	json.NewEncoder(w).Encode(restaurants)
}

func GetAllDishesHandler(w http.ResponseWriter, r *http.Request) {
	dishes, err := dbHelper.GetAllDishes()
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
