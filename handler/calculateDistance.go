package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/middleware"
	"rmssystem_1/services"
	"rmssystem_1/utils"
)

func CalculateDistance(w http.ResponseWriter, r *http.Request) {
	_, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	restId := r.URL.Query().Get("rest_id")
	addId := r.URL.Query().Get("add_id")

	if restId == "" || addId == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "restaurant_id and address_id are required query parameters")
		return
	}

	distanceKm, err := services.CalculateRestaurantDistanceByID(restId, addId)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to calculate distance")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "successfully calculated distance",
		"distance between house and restaurant is (in km) ::": distanceKm,
	})
}
