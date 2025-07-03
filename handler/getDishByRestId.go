package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/utils"
)

func GetDishesByRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID := r.URL.Query().Get("restaurant_id")
	limit, offset := utils.GetPageLimitAndOffset(r)
	if restaurantID == "" {
		w.WriteHeader(http.StatusBadRequest)
		jsoniter.NewEncoder(w).Encode(map[string]interface{}{
			"message": "restaurant_id is required",
		})
		return
	}
	dishes, err := dbHelper.GetDishesByRestaurant(restaurantID, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsoniter.NewEncoder(w).Encode(map[string]interface{}{
			"message": "failed to fetch dishes",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dishes fetched successfully",
		"data":    dishes,
	})
}
