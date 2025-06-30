package handler

import (
	jsoniter "github.com/json-iterator/go"
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
	jsoniter.NewEncoder(w).Encode(restaurants)
}

func GetAllDishesHandler(w http.ResponseWriter, r *http.Request) {
	dishes, err := dbHelper.GetAllDishes()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch dishes")
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dishes fetched successfully",
		"data":    dishes,
	})
}
