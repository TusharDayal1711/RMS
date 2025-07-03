package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/utils"
)

// getting all the restaurent form db
func ListAllRestaurants(w http.ResponseWriter, r *http.Request) {
	limit, offset := utils.GetPageLimitAndOffset(r)
	restaurants, err := dbHelper.GetAllRestaurant(limit, offset)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch restaurants")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "List Of All Restaurants",
		"Restaurants": restaurants,
	})
}

func GetAllDishesHandler(w http.ResponseWriter, r *http.Request) {

	limit, offset := utils.GetPageLimitAndOffset(r)
	dishes, err := dbHelper.GetAllDishes(limit, offset)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch dishes")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "List Of Dishes",
		"Dishes":  dishes,
	})
}
